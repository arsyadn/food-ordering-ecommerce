package repositories

import (
	"food-ordering/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order models.Order) (models.Order, error)
	FindByUserID(userID uint) ([]models.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) Create(order models.Order) (models.Order, error) {
	err := r.db.Create(&order).Error
	return order, err
}

func (r *orderRepository) FindByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("Items").Where("user_id = ?", userID).Order("created_at desc").Find(&orders).Error
	return orders, err
}
