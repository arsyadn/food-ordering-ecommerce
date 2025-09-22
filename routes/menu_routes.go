package routes

import (
	"food-ordering/controllers"
	"food-ordering/middleware"
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
	menuRoutes.Use(middleware.AuthMiddleware())
	menuRoutes.Use(middleware.RoleAdminMiddleware())

	userRoutes := rg.Group("/menus")
	userRoutes.Use(middleware.AuthMiddleware())

	{
		userRoutes.GET("/", menuController.GetMenus)
		userRoutes.GET("/:id", menuController.GetMenuByID)
		menuRoutes.POST("/", menuController.CreateMenu)
		menuRoutes.PUT("/:id", menuController.UpdateMenu)
		menuRoutes.DELETE("/:id", menuController.DeleteMenu)
	}
}
