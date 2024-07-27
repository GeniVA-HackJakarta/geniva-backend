package repository

import "gorm.io/gorm"

var (
	UserRepo       UserRepository
	MenuRepo       MenuRepository
	RestaurantRepo RestaurantRepository
	HistoryRepo    HistoryRepository
	OrderRepo      OrderRepository
	DriverRepo     DriverRepository
)

func InitRepository(db *gorm.DB) {
	UserRepo = NewUserRepository(db)
	MenuRepo = NewMenuRepository(db)
	RestaurantRepo = NewRestaurantRepository(db)
	HistoryRepo = NewHistoryRepository(db)
	OrderRepo = NewOrderRepository(db)
	DriverRepo = NewDriverRepository(db)
}
