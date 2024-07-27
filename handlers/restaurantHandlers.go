package handlers

import (
	"geniva/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllRestaurants(c *gin.Context) {
	restaurants, err := repository.RestaurantRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to retrieve restaurants", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "restaurants": restaurants})
}

func GetRestaurant(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid ID"})
		return
	}
	menu, err := repository.MenuRepo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to retrieve restaurant", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "menu": menu})
}
