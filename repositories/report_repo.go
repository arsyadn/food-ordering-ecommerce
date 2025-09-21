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
	err := r.db.Preload("User").Preload("Items.Menu").Find(&orders).Error
	return orders, err
}
