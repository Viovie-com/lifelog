package api

import (
	"net/http"
	"strconv"

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
		post.PostTags = append(post.PostTags, tag)
	}

	post.Create()
}

type tagResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type postResponse struct {
	ID         uint          `json:"id"`
	MemberID   uint          `json:"memberId"`
	Title      string        `json:"title"`
	Content    string        `json:"content"`
	CategoryID int           `json:"categoryId"`
	Draft      bool          `json:"draft"`
	Tags       []tagResponse `jsons:"tags"`
}

func GetPosts(c *gin.Context) {
	request := map[string]string{
		"page":  c.DefaultQuery("page", "1"),
		"limit": c.DefaultQuery("limit", "10"),
	}
	page, err := strconv.ParseInt(request["page"], 10, 32)
	if err != nil {
		page = 1
	}

	limit, err := strconv.ParseInt(request["limit"], 10, 32)
	if err != nil {
		limit = 10
	}

	posts := db.GetPosts(int(page), int(limit))

	var response []postResponse
	for _, post := range posts {
		var tags []tagResponse
		for _, tag := range post.PostTags {
			tags = append(tags, tagResponse{
				ID:   tag.TagID,
				Name: tag.Tag.Name,
			})
		}

		response = append(response, postResponse{
			ID:         post.ID,
			MemberID:   post.MemberID,
			Title:      post.Title,
			Content:    post.Content,
			CategoryID: post.CategoryID,
			Draft:      post.Draft,
			Tags:       tags,
		})
	}

	if response == nil {
		response = []postResponse{}
	}

	c.JSON(http.StatusOK, response)
}
