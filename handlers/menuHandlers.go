package handlers

import (
	"geniva/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllMenus(c *gin.Context) {
	menus, err := repository.MenuRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to retrieve menus", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "menus": menus})
}

func GetMenu(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid ID"})
		return
	}
	menu, err := repository.MenuRepo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to retrieve menu", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "menu": menu})
}
