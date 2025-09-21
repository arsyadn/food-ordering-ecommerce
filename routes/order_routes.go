package routes

import (
	"food-ordering/controllers"
	"food-ordering/middleware"
	"food-ordering/repositories"
	"food-ordering/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupOrderRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	orderRepo := repositories.NewOrderRepository(db)
	cartRepo := repositories.NewCartRepository(db)
	menuRepo := repositories.NewMenuRepository(db)
	orderService := services.NewOrderService(orderRepo, cartRepo, menuRepo, db)
	orderController := controllers.NewOrderController(orderService)

	orderRoutes := rg.Group("/")
	orderRoutes.Use(middleware.AuthMiddleware())
	{
		orderRoutes.POST("/checkout", orderController.Checkout)
		orderRoutes.GET("/orders", orderController.GetOrders)
	}
}
