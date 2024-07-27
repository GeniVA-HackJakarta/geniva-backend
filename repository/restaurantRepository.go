package repository

import (
	"geniva/models"

	"gorm.io/gorm"
)

type RestaurantRepository interface {
	GetAll() ([]models.Restaurant, error)
	GetByID(id int) (models.Restaurant, error)
}

type restaurantRepo struct {
	db *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) RestaurantRepository {
	return &restaurantRepo{db}
}

func (repo *restaurantRepo) GetAll() ([]models.Restaurant, error) {
	var restaurants []models.Restaurant
	result := repo.db.Find(&restaurants)
	return restaurants, result.Error
}

func (repo *restaurantRepo) GetByID(id int) (models.Restaurant, error) {
	var restaurant models.Restaurant
	result := repo.db.First(&restaurant, id)
	return restaurant, result.Error
}
