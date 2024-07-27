package routes

import (
	"geniva/handlers"

	"github.com/gin-gonic/gin"
)

func ConfigureRouter(router *gin.Engine) {
	router.GET("/api", handlers.MainHandler)
	// apis := router.Group("/api")
}
