package repository

import (
	"geniva/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *models.Order) error
	GetOrderById(orderId int) (*models.Order, error)
	GetOrdersByUserId(userId int) ([]models.Order, error)
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepo{db}
}

func (repo *orderRepo) CreateOrder(order *models.Order) error {
	result := repo.db.Create(order)
	return result.Error
}

func (repo *orderRepo) GetOrderById(orderId int) (*models.Order, error) {
	var order models.Order
	result := repo.db.First(&order, orderId)
	return &order, result.Error
}

func (repo *orderRepo) GetOrdersByUserId(userId int) ([]models.Order, error) {
	var orders []models.Order
	result := repo.db.Where("user_id = ?", userId).Find(&orders)
	return orders, result.Error
}
