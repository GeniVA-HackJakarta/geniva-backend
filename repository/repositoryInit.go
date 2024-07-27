package repository

import "gorm.io/gorm"

var (
	UserRepo UserRepository
)

func InitRepository(db *gorm.DB) {
	UserRepo = NewUserRepository(db)
}
