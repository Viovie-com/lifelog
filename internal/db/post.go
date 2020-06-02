package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	MemberID   uint   `gorm:"column:member_id"`
	Title      string `gorm:"column:title"`
	Content    string `gorm:"column:content"`
	CategoryID int    `gorm:"column:category_id"`
	Draft      bool   `gorm:"column:draft"`
	PostTags   []*PostTag
}

func (Post) TableName() string {
	return "post"
}

type Tag struct {
	Model
	Name     string `gorm:"column:name"`
	PostTags []*PostTag
}

func (Tag) TableName() string {
	return "tag"
}

type PostTag struct {
	PostID    uint `gorm:"column:post_id"`
	Post      *Post
	TagID     uint `gorm:"column:tag_id"`
	Tag       *Tag
	CreatedAt time.Time
}

func (PostTag) TableName() string {
	return "post_tag"
}

func (post *Post) Create() (err error) {
	db := Instance()
	defer db.Close()

	for _, postTag := range post.PostTags {
		db.Where(postTag.Tag).FirstOrInit(postTag.Tag)
	}

	db.Create(post)

	return
}
