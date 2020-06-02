package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Viovie-com/lifelog/internal/db"
)

type registerRequest struct {
	Name     string  `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Avatar   *string `json:"avatar"`
	Password string  `json:"password" binding:"required"`
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
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	member := db.GetMember(uint(id))

	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	member.Update()

	c.JSON(http.StatusOK, gin.H{})
}
