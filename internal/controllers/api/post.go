package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Viovie-com/lifelog/internal/db"
	"github.com/Viovie-com/lifelog/internal/middlewares"
)

type addPostRequest struct {
	Title      string   `json:"title" binding:"required"`
	Content    string   `json:"content" binding:"required"`
	CategoryID int      `json:"categoryId" binding:"required,number"`
	Draft      bool     `json:"draft" binding:"required"`
	Tags       []string `json:"tags" binding:"required"`
}

func AddPost(c *gin.Context) {
	member, err := middlewares.GetMemberFromAuth(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var request addPostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	post := db.Post{
		MemberID:   member.ID,
		Title:      request.Title,
		Content:    request.Content,
		CategoryID: request.CategoryID,
		Draft:      request.Draft,
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
