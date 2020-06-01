package server

import (
	"github.com/gin-gonic/gin"

	"github.com/Viovie-com/lifelog/internal/controllers"
)

func registerRoutes(router *gin.Engine) {

	// Set template
	router.LoadHTMLGlob("web/template/*")

	// Set asserts
	router.Static("js", "web/public/js")
	router.Static("css", "web/public/css")
	router.Static("images", "web/public/images")

	router.GET("/", controllers.GetIndex)

	// Api group
	api := router.Group("/api/v1")

	// Member api
	memberApi := api.Group("/member")
	{
		memberApi.POST("/", controllers.RegisterMember)
		memberApi.PUT("/:id", controllers.UpdateMember)
	}

	authApi := api.Group("/auth")
	{
		authApi.POST("/", controllers.Login)
	}
}
