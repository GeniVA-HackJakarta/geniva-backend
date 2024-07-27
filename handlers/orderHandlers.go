package handlers

import (
	"geniva/models"
	"geniva/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateOrderRequest struct {
	UserID         int     `json:"user_id"`
	DriverID       int     `json:"driver_id"`
	VehicleID      int     `json:"vehicle_id"`
	FoodItemID     int     `json:"food_item_id,omitempty"`
	TotalPrice     float64 `json:"total_price"`
	Discount       float64 `json:"discount"`
	DiscountName   string  `json:"discount_name"`
	FinalPrice     float64 `json:"final_price"`
	Status         string  `json:"status"`
	OrderType      string  `json:"order_type"` // "bike", "car", "food"
	PickupLocation string  `json:"pickup_location"`
	Destination    string  `json:"destination"`
}

func CreateOrder(c *gin.Context) {
	var req CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	order := models.Order{
		UserID:       req.UserID,
		DriverID:     req.DriverID,
		VehicleID:    req.VehicleID,
		FoodItemID:   req.FoodItemID,
		TotalPrice:   req.TotalPrice,
		Discount:     req.Discount,
		DiscountName: req.DiscountName,
		FinalPrice:   req.FinalPrice,
		Status:       req.Status,
		OrderType:    req.OrderType,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := repository.OrderRepo.CreateOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	history := models.History{
		UserID:         req.UserID,
		ServiceType:    req.OrderType,
		OrderTime:      time.Now(),
		PickupLocation: req.PickupLocation,
		Destination:    req.Destination,
		Fare:           req.FinalPrice,
		Status:         req.Status,
	}

	err = repository.HistoryRepo.SaveHistory(&history)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order history"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "order_id": order.ID})
}

func GetOrderHistory(c *gin.Context) {
	userIDParam := c.Param("user_id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	orders, err := repository.OrderRepo.GetOrdersByUserId(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get order history"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "orders": orders})
}
