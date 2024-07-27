package handlers

import (
	"geniva/models"
	"geniva/repository"
	"geniva/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid user data"})
		return
	}

	storedUser, err := repository.UserRepo.GetUserByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Invalid user email"})
		return
	}

	if err := utils.CheckHashedPassword(user.Password, storedUser.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Invalid user password"})
		return
	}

	tokenString, err := utils.CreateToken(uint(storedUser.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to create token"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"success": true,
		"id":      storedUser.ID,
		"message": "Welcome, " + storedUser.Name + "!",
		"token":   tokenString,
	})
}

func Register(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid user data"})
		return
	}

	existingUser, err := repository.UserRepo.GetUserByEmail(newUser.Email)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"success": false, "error": "User " + existingUser.Name + " already exists"})
		return
	}

	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to hash password"})
		return
	}
	newUser.Password = string(hashedPassword)
	newUser.KYC = false

	err = repository.UserRepo.SaveUser(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to register user"})
		return
	}

	tokenString, err := utils.CreateToken(uint(newUser.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to create token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"id":      newUser.ID,
		"message": "Welcome, " + newUser.Name + "!",
		"token":   tokenString,
	})

}
