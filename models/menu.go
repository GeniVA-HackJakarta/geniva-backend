package models

type Menu struct {
	MenuID        string `json:"menu_id" gorm:"primaryKey"`
	MenuName      string `json:"menu_name"`
	RestaurantID  string `json:"restaurant_id"`
	Price         string `json:"price"`
	GlucoseLevels string `json:"glucose_levels"`
}
