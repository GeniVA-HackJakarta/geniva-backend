package repository

import (
	"geniva/models"
	"time"

	"math/rand"

	"gorm.io/gorm"
)

type DriverRepository interface {
	GetRandomDriver() (*models.Driver, error)
}

type driverRepo struct {
	db *gorm.DB
}

func NewDriverRepository(db *gorm.DB) DriverRepository {
	return &driverRepo{db}
}

func (repo *driverRepo) GetRandomDriver() (*models.Driver, error) {
	var drivers []models.Driver
	result := repo.db.Find(&drivers)
	if result.Error != nil {
		return nil, result.Error
	}

	if len(drivers) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(drivers))
	return &drivers[randomIndex], nil
}
