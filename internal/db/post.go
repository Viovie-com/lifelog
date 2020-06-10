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
	PostTags   []PostTag
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

type tmpTag struct {
	PostID uint
	ID     uint
	Name   string
}

func GetPosts(page int, limit int) (posts []Post) {
	db := Instance()
	defer db.Close()

	db.Offset((page - 1) * limit).Limit(limit).Find(&posts)

	var postIds []uint
	for _, post := range posts {
		postIds = append(postIds, post.ID)
	}

	tagMap := make(map[uint][]PostTag)
	rows, _ := db.Table("post_tag").Select("post_tag.post_id, tag.id, tag.name").Joins("INNER JOIN tag ON post_tag.tag_id = tag.id").Where("post_tag.post_id IN (?)", postIds).Rows()
	for rows != nil && rows.Next() {
		var tmp tmpTag
		db.ScanRows(rows, &tmp)
		tagMap[tmp.PostID] = append(tagMap[tmp.PostID], PostTag{
			PostID: tmp.PostID,
			TagID:  tmp.ID,
			Tag: &Tag{
				Name: tmp.Name,
			},
		})
	}

	for k, post := range posts {
		posts[k].PostTags = tagMap[post.ID]
	}
	return
}
