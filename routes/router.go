package routes

import (
	"geniva/handlers"

	"github.com/gin-gonic/gin"
)

func ConfigureRouter(router *gin.Engine) {
	router.GET("/api", handlers.MainHandler)
	apis := router.Group("/api")

	// Auth Controller
	apis.POST("/auth/register", handlers.Register)
	apis.POST("/auth/login", handlers.Login)

	// User Handler
	apis.GET("/user", handlers.GetAllUsers)
	apis.GET("/user/:id", handlers.GetUser)
	apis.PUT("/user/:id/kyc", handlers.UpdateKYC)

	// Menu Handler
	apis.GET("/menu", handlers.GetAllMenus)
	apis.GET("/menu/:id", handlers.GetMenu)

	// Restaurant Handler
	apis.GET("/restaurant", handlers.GetAllRestaurants)
	apis.GET("/restaurant/:id", handlers.GetRestaurant)

	// History Handler 
	apis.GET("/history/:user_id", handlers.GetHistoryByUserID)
	apis.POST("/history", handlers.SaveHistory)

	// Order Handler
	apis.POST("/order", handlers.CreateOrder)

	// AI Handler
	apis.POST("/ai/:user_id", handlers.CallAIHandler)

	// Final Price + Saving Handler
	apis.POST("/price/calculate", handlers.CalculateFinalPrice)
}
