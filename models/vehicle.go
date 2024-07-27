package models

type Vehicle struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Type string `json:"type"`
}

type Driver struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	VehicleID int    `json:"vehicle_id"`
}
