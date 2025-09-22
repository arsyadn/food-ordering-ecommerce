package routes

import (
	"food-ordering/controllers"
	"food-ordering/middleware"
	"food-ordering/repositories"
	"food-ordering/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupCartRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	cartRepo := repositories.NewCartRepository(db)
	cartService := services.NewCartService(cartRepo)
	cartController := controllers.NewCartController(cartService)

	cartRoutes := rg.Group("/cart")
	cartRoutes.Use(middleware.AuthMiddleware())

	{
		cartRoutes.GET("/", cartController.GetCart)
		cartRoutes.POST("/:id", cartController.AddToCart)
		cartRoutes.PUT("/:id", cartController.UpdateCart)
		cartRoutes.DELETE("/:id", cartController.DeleteCart)
	}
}
