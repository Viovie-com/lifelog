package server

import (
	"github.com/gin-gonic/gin"

	"github.com/Viovie-com/lifelog/internal/controllers"
)

func registerRoutes(router *gin.Engine) {
	// TODO: change path
	basePath := "src/lifelog/"

	// Set template
	router.LoadHTMLGlob(basePath + "web/template/*")

	// Set asserts
	router.Static("js", basePath + "web/public/js")
	router.Static("css", basePath + "web/public/css")
	router.Static("images", basePath + "web/public/images")

	router.GET("/", controllers.GetIndex)
}
