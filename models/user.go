package models

type User struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Email       string `json:"email" gorm:"unique"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	KYC         bool   `json:"kyc"`
	Saving      string `json:"saving"`
}
