package server

import (
	"github.com/gin-gonic/gin"

	"github.com/Viovie-com/lifelog/internal/controllers"
	controllerApi "github.com/Viovie-com/lifelog/internal/controllers/api"
	"github.com/Viovie-com/lifelog/internal/middlewares"
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
		memberApi.POST("/", controllerApi.RegisterMember)

		memberApi.Use(middlewares.Auth())
		memberApi.PUT("/", controllerApi.UpdateMember)
	}

	// Auth api
	authApi := api.Group("/auth")
	{
		authApi.POST("/", controllerApi.Login)
	}

	// Post api
	postApi := api.Group("/post")
	{
		postApi.GET("/", controllerApi.GetPosts)

		postApi.Use(middlewares.Auth())
		postApi.POST("/", controllerApi.AddPost)
	}
}
