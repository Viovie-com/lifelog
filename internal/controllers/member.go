package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Viovie-com/lifelog/internal/db"
)

func RegisterMember(c *gin.Context) {
	var member db.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
