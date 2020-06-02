package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Viovie-com/lifelog/internal/db"
)

type addPostRequest struct {
	MemberID   uint     `json:"memberId" binding:"required"`
	Title      string   `json:"title" binding:"required"`
	Content    string   `json:"content" binding:"required"`
	CategoryID int      `json:"categoryId" binding:"required"`
	Draft      bool     `json:"draft" binding:"required"`
	Tags       []string `json:"tags" binding:"required"`
}

func AddPost(c *gin.Context) {
	var request addPostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := db.Post{
		MemberID: request.MemberID,
		Title: request.Title,
		Content: request.Content,
		CategoryID: request.CategoryID,
		Draft: request.Draft,
	}

	for _, name := range request.Tags {
		tag := db.PostTag{
			Tag: &db.Tag{
				Name: name,
			},
		}
		post.PostTags = append(post.PostTags, &tag)
	}

	post.Create()
}
