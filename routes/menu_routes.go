package routes

import (
	"food-ordering/controllers"
	"food-ordering/repositories"
	"food-ordering/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupMenuRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	menuRepo := repositories.NewMenuRepository(db)
	menuService := services.NewMenuService(menuRepo)
	menuController := controllers.NewMenuController(menuService)

	menuRoutes := rg.Group("/menus")
	{
		menuRoutes.GET("/", menuController.GetMenus)
		menuRoutes.GET("/:id", menuController.GetMenuByID)
		menuRoutes.POST("/", menuController.CreateMenu)
	}
}
