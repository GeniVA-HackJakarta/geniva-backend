package models

import "github.com/google/uuid"

type User struct {
	UUID     uuid.UUID `json:"uuid" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name     string    `json:"name"`
	Email    string    `json:"email" gorm:"unque"`
	Password string    `json:"password"`
}
