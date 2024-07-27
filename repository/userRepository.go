package repository

import (
	"geniva/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

func (repo *userRepo) GetAll() ([]models.User, error) {
	var users []models.User
	result := repo.db.Find(&users)
	return users, result.Error
}
