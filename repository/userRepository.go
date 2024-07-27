package repository

import (
	"geniva/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(id int) (models.User, error)
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

func (repo *userRepo) GetByID(id int) (models.User, error) {
	var user models.User
	result := repo.db.First(&user, id)
	return user, result.Error
}

