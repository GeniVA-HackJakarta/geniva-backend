package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, Welcome to GeniVA Web Service!")
}
