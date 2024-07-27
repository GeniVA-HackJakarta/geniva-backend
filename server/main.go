package main

import (
	"geniva/config"
	middlewares "geniva/middlewares"
	"geniva/routes"
	initializers "geniva/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	router := gin.Default()
	router.Use(middlewares.CorsMiddleware())
	config.Connect()
	routes.ConfigureRouter(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
