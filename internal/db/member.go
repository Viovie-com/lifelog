package db

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

type AuthSource int

const (
	AuthEmail AuthSource = iota
	AuthFacebook
	AuthGoogle
)

func (source AuthSource) String() string {
	return [3]string{"Email", "Facebook", "Google"}[source]
}

type Member struct {
	gorm.Model
	Name   string      `gorm:"column:name" json:"name" binding:"required"`
	Email  *string     `gorm:"column:email" json:"email"`
	Avatar *string     `gorm:"column:avatar" json:"avatar"`
	Auth   *MemberAuth `gorm:"-"`
}

func (Member) TableName() string {
	return "member"
}

type MemberAuth struct {
	Model
	MemberID     uint      `gorm:"column:member_id"`
	Source       string    `gorm:"column:source"`
	SourceID     string    `gorm:"column:source_id"`
	SourceToken  string    `gorm:"column:source_token"`
	Token        string    `gorm:"column:token"`
	RefreshToken string    `gorm:"column:refresh_token"`
	ExpiredAt    time.Time `gorm:"expired_at"`
}

func (MemberAuth) TableName() string {
	return "member_auth"
}

func (member *Member) Register() error {
	db := Instance()
	defer db.Close()

	transaction := db.Begin()
	if err := transaction.Create(member).Error; err != nil {
		transaction.Rollback()
		return err
	}

	member.Auth.MemberID = member.ID
	if err := transaction.Create(member.Auth).Error; err != nil {
		transaction.Rollback()
		return err
	}

	return transaction.Commit().Error
}

func (member *Member) Update() {
	db := Instance()
	defer db.Close()

	db.Save(member)
}

func GetMember(id uint) (member Member) {
	db := Instance()
	defer db.Close()

	db.Where(id).First(&member)
	return
}

func GetMemberByToken(token string) (member Member, err error) {
	db := Instance()
	defer db.Close()

	var memberAuth MemberAuth
	err = db.Where(MemberAuth{Token: token}).First(&memberAuth).Error
	if err != nil {
		return
	}
	err = db.Where(memberAuth.MemberID).First(&member).Error
	return
}

func Encrypt(password string) (token string) {
	h := sha256.New()
	h.Write([]byte(password))
	token = hex.EncodeToString(h.Sum(nil))
	return
}

func Login(email string, password string) (auth MemberAuth, err error) {
	db := Instance()
	defer db.Close()

	password = Encrypt(password)

	db = db.Where(&MemberAuth{Source: AuthEmail.String(), SourceID: email, SourceToken: password}).First(&auth)
	if err = db.Error; err != nil {
		fmt.Println(err)
		return
	}

	auth.Create()
	return
}

func (auth *MemberAuth) Create() {
	auth.ExpiredAt = time.Now().AddDate(0, 0, 15)
	auth.Token = Encrypt(auth.SourceID + strconv.FormatInt(auth.ExpiredAt.Unix(), 10))
	auth.RefreshToken = Encrypt(strconv.FormatUint(uint64(auth.ID), 10) + auth.SourceID + strconv.FormatInt(auth.ExpiredAt.Unix(), 10))

	db := Instance()
	defer db.Close()

	db.Save(auth)
}
