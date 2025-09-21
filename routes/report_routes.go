package routes

import (
	"food-ordering/controllers"
	"food-ordering/repositories"
	"food-ordering/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupReportRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	reportRepo := repositories.NewReportRepository(db)
	reportService := services.NewReportService(reportRepo)
	reportController := controllers.NewReportController(reportService)

	reports := rg.Group("/reports")
	{
		reports.GET("/sales", reportController.GetSalesReport)
	}
}
