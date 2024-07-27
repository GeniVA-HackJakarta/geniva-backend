package models

import "time"

type Order struct {
	ID           int       `json:"id" gorm:"primary_key"`
	UserID       int       `json:"user_id"`
	DriverID     int       `json:"driver_id"`
	VehicleID    int       `json:"vehicle_id"`
	FoodItemID   int       `json:"food_item_id,omitempty"`
	TotalPrice   float64   `json:"total_price"`
	Discount     float64   `json:"discount"`
	DiscountName string    `json:"discount_name"`
	FinalPrice   float64   `json:"final_price"`
	Status       string    `json:"status"`
	OrderType    string    `json:"order_type"` // "bike", "car", "food"
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
