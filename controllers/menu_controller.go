package controllers

import (
	"food-ordering/models"
	"food-ordering/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
	menuService services.MenuService
}

func NewMenuController(menuService services.MenuService) *MenuController {
	return &MenuController{menuService}
}

func (c *MenuController) GetMenus(ctx *gin.Context) {
	menus, err := c.menuService.GetAllMenus()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, menus)
}

func (c *MenuController) GetMenuByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	menu, err := c.menuService.GetMenuByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "menu not found"})
		return
	}
	ctx.JSON(http.StatusOK, menu)
}

func (c *MenuController) CreateMenu(ctx *gin.Context) {
	var menu models.Menu
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdMenu, err := c.menuService.CreateMenu(menu)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, createdMenu)
}
