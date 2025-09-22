package routes

import (
	"food-ordering/controllers"
	"food-ordering/middleware"
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
	reports.Use(middleware.AuthMiddleware())
	reports.Use(middleware.RoleAdminMiddleware())
	{
		reports.GET("/sales", reportController.GetSalesReport)
	}
}
