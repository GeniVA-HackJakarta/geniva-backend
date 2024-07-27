package handlers

import (
	"geniva/models"
	"geniva/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveHistory(c *gin.Context) {
	var request struct {
		UserID         int     `json:"user_id"`
		ServiceType    string  `json:"service_type"`
		PickupLocation string  `json:"pickup_location"`
		Destination    string  `json:"destination"`
		Fare           float64 `json:"fare"`
		Status         string  `json:"status"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid request"})
		return
	}

	history := models.History{
		UserID:         request.UserID,
		ServiceType:    request.ServiceType,
		OrderTime:      time.Now(),
		PickupLocation: request.PickupLocation,
		Destination:    request.Destination,
		Fare:           request.Fare,
		Status:         request.Status,
	}

	if err := repository.HistoryRepo.SaveHistory(&history); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to save history", "detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "History saved successfully"})
}

func GetHistoryByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	history, err := repository.HistoryRepo.GetHistoryByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve history"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"history": history})
}
