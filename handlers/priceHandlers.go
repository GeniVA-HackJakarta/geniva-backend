package handlers

import (
	"geniva/repository"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PriceRequest struct {
	UserID int     `json:"user_id"`
	Price  float64 `json:"price"`
}

func CalculateFinalPrice(c *gin.Context) {
	var request PriceRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := repository.UserRepo.GetUserByID(request.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	originalPrice := request.Price
	roundedPrice := math.Ceil(originalPrice/1000) * 1000 // Round up to nearest thousand
	finalPrice := roundedPrice
	var savings float64

	if roundedPrice <= 20000 {
		switch user.Saving {
		case "MEDIUM":
			finalPrice += 1000
		case "HIGH":
			finalPrice += 2000
		}
	} else {
		switch user.Saving {
		case "MEDIUM":
			finalPrice += 2000
		case "HIGH":
			finalPrice += 4000
		}
	}

	savings = finalPrice - originalPrice
	savings = math.Round(savings*100) / 100 // Round savings to 2 decimal places

	c.JSON(http.StatusOK, gin.H{
		"original_price": originalPrice,
		"rounded_price":  roundedPrice,
		"final_price":    finalPrice,
		"savings":        savings,
	})
}
