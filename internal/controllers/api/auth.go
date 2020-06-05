package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Viovie-com/lifelog/internal/db"
)

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var request loginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auth, err := db.Login(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": auth.Token, "refreshToken": auth.RefreshToken, "expiredAt": auth.ExpiredAt})
}
