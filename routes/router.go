package routes

import (
	"geniva/handlers"

	"github.com/gin-gonic/gin"
)

func ConfigureRouter(router *gin.Engine) {
	router.GET("/api", handlers.MainHandler)
	apis := router.Group("/api")

	// User Handler
	apis.GET("/user", handlers.GetAllUsers)
	apis.GET("/user/:id", handlers.GetUser)

	// Menu Handler
	apis.GET("/menu", handlers.GetAllMenus)
	apis.GET("/menu/:id", handlers.GetMenu)

	// Restaurant Handler
	apis.GET("/restaurant", handlers.GetAllRestaurants)
	apis.GET("/restaurant/:id", handlers.GetRestaurant)

}
