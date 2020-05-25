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
}
