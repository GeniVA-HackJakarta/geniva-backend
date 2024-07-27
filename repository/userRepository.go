package repository

import (
	"geniva/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(id int) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	SaveUser(user *models.User) error
	UpdateKYC(id int, saving string) error
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

func (repo *userRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	result := repo.db.Where("email = ?", email).First(&user)
	return user, result.Error
}

func (repo *userRepo) SaveUser(user *models.User) error {
	result := repo.db.Create(user)
	return result.Error
}

func (repo *userRepo) UpdateKYC(id int, saving string) error {
	result := repo.db.Model(&models.User{}).Where("id = ?", id).Updates(map[string]interface{}{"kyc": true, "saving": saving})
	return result.Error
}
