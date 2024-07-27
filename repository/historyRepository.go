package repository

import (
	"geniva/models"

	"gorm.io/gorm"
)

type HistoryRepository interface {
	SaveHistory(history *models.History) error
	GetHistoryByUserID(userID int) ([]models.History, error)
}

type historyRepo struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) HistoryRepository {
	return &historyRepo{db}
}

func (repo *historyRepo) SaveHistory(history *models.History) error {
	return repo.db.Create(history).Error
}

func (repo *historyRepo) GetHistoryByUserID(userID int) ([]models.History, error) {
	var history []models.History
	err := repo.db.Where("user_id = ?", userID).Find(&history).Error
	return history, err
}
