package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Viovie-com/lifelog/internal/db"
	"github.com/Viovie-com/lifelog/internal/middlewares"
)

type registerRequest struct {
	Name     string  `json:"name" binding:"required,max=100"`
	Email    string  `json:"email" binding:"required,email"`
	Avatar   *string `json:"avatar"`
	Password string  `json:"password" binding:"required"`
}

type updateRequest struct {
	Name   *string `gorm:"column:name" json:"name"`
	Avatar *string `gorm:"column:avatar" json:"avatar"`
}

func RegisterMember(c *gin.Context) {
	var request registerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auth := db.MemberAuth{
		Source:      db.AuthEmail.String(),
		SourceID:    request.Email,
		SourceToken: db.Encrypt(request.Password),
	}

	member := db.Member{
		Name:   request.Name,
		Email:  &request.Email,
		Avatar: request.Avatar,
		Auth:   &auth,
	}

	member.Register()

	c.JSON(http.StatusOK, gin.H{})
}

func UpdateMember(c *gin.Context) {
	member, err := middlewares.GetMemberFromAuth(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var request updateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if request.Name != nil {
		member.Name = *request.Name
	}
	if request.Avatar != nil {
		member.Avatar = request.Avatar
	}
	member.Update()

	c.JSON(http.StatusOK, gin.H{})
}
