package db

import (
	"crypto/sha256"
	"encoding/hex"

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
	MemberID    uint   `gorm:"column:member_id"`
	Source      string `gorm:"column:source"`
	SourceID    string `gorm:"column:source_id"`
	SourceToken string `gorm:"column:source_token"`
	Token       string `gorm:"column:token"`
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

func Encrypt(password string) (token string) {
	h := sha256.New()
	h.Write([]byte(password))
	token = hex.EncodeToString(h.Sum(nil))
	return
}
