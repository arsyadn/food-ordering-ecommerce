package repositories

import (
	"food-ordering/models"

	"gorm.io/gorm"
)

type ReportRepository interface {
	GetSalesReport() ([]models.Order, error)
}

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{db}
}

func (r *reportRepository) GetSalesReport() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Raw(`
		SELECT o.* 
		FROM orders o
		JOIN order_items oi ON oi.order_id = o.id
		JOIN menus m ON oi.menu_id = m.id
	`).Scan(&orders).Error
	return orders, err
}
