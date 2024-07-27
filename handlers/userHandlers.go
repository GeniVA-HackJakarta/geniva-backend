package handlers

import (
	"geniva/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users, err := repository.UserRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to retrieve users", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "users": users})
}

func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid ID"})
		return
	}
	user, err := repository.UserRepo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to retrieve user", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "user": user})
}
