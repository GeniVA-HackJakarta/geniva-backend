package models

import "time"

type History struct {
	ID             int       `json:"id" gorm:"primaryKey"`
	UserID         int       `json:"user_id"`
	ServiceType    string    `json:"service_type"`
	OrderTime      time.Time `json:"order_time"`
	PickupLocation string    `json:"pickup_location"`
	Destination    string    `json:"destination"`
	Fare           float64   `json:"fare"`
	Status         string    `json:"status"`
	FoodItemID     int      `json:"food_item_id,omitempty"`
}
