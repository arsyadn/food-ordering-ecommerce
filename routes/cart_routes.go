package routes

import (
	"food-ordering/controllers"
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
	{
		cartRoutes.GET("/", cartController.GetCart)
		cartRoutes.POST("/", cartController.AddToCart)
		cartRoutes.PUT("/:id", cartController.UpdateCart)
		cartRoutes.DELETE("/:id", cartController.DeleteCart)
	}
}
