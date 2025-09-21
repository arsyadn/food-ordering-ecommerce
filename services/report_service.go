package services

import (
	"food-ordering/models"
	"food-ordering/repositories"
)

type ReportService interface {
	GetSalesReport() ([]models.Order, error)
}

type reportService struct {
	reportRepo repositories.ReportRepository
}

func NewReportService(repo repositories.ReportRepository) ReportService {
	return &reportService{reportRepo: repo}
}

func (s *reportService) GetSalesReport() ([]models.Order, error) {
	return s.reportRepo.GetSalesReport()
}
