package db

import (
	"github.com/jinzhu/gorm"
)

type Member struct {
	gorm.Model
	Name   string  `gorm:"column:name" json:"name" binding:"required"`
	Email  *string `gorm:"column:email" json:"email"`
	Avatar *string `gorm:"column:avatar" json:"avatar"`
}

func (Member) TableName() string {
	return "member"
}

func (member *Member) Register() {
	db := Instance()
	defer db.Close()

	db.Create(member)
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
