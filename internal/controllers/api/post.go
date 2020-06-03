package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Viovie-com/lifelog/internal/db"
)

type addPostRequest struct {
	Title      string   `json:"title" binding:"required"`
	Content    string   `json:"content" binding:"required"`
	CategoryID int      `json:"categoryId" binding:"required"`
	Draft      bool     `json:"draft" binding:"required"`
	Tags       []string `json:"tags" binding:"required"`
}

func AddPost(c *gin.Context) {
	var request addPostRequest
	m, hasAuth := c.Get("member")
	if err := c.ShouldBindJSON(&request); err != nil || !hasAuth {
		if !hasAuth {
			err = errors.New("access is denied (5)")
		}
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	member, ok := m.(db.Member)
	if !ok {
		err := errors.New("access is denied (6)")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	post := db.Post{
		MemberID: member.ID,
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
