package repository

import "gorm.io/gorm"

var (
	UserRepo UserRepository
	MenuRepo MenuRepository
	RestaurantRepo RestaurantRepository
)

func InitRepository(db *gorm.DB) {
	UserRepo = NewUserRepository(db)
	MenuRepo = NewMenuRepository(db)
	RestaurantRepo = NewRestaurantRepository(db)
}
