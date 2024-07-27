package repository

import (
	"geniva/models"

	"gorm.io/gorm"
)

type MenuRepository interface {
	GetAll() ([]models.Menu, error)
	GetByID(id int) (models.Menu, error)
}

type menuRepo struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepo{db}
}

func (repo *menuRepo) GetAll() ([]models.Menu, error) {
	var menus []models.Menu
	result := repo.db.Find(&menus)
	return menus, result.Error
}

func (repo *menuRepo) GetByID(id int) (models.Menu, error) {
	var menu models.Menu
	result := repo.db.First(&menu, id)
	return menu, result.Error
}
