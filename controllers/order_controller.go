package controllers

import (
	"food-ordering/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(orderService services.OrderService) *OrderController {
	return &OrderController{orderService}
}

// POST /checkout
func (c *OrderController) Checkout(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(uint)

	order, err := c.orderService.Checkout(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, order)
}

// GET /orders
func (c *OrderController) GetOrders(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(uint)

	orders, err := c.orderService.GetUserOrders(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, orders)
}
