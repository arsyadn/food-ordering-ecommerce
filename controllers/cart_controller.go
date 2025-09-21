package controllers

import (
	"food-ordering/models"
	"food-ordering/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartController struct {
	cartService services.CartService
}

func NewCartController(cartService services.CartService) *CartController {
	return &CartController{cartService}
}

// GET /cart
func (c *CartController) GetCart(ctx *gin.Context) {
	userIDStr := ctx.Query("user_id")
	userID, _ := strconv.Atoi(userIDStr)

	carts, err := c.cartService.GetUserCart(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, carts)
}

// POST /cart
func (c *CartController) AddToCart(ctx *gin.Context) {
	var cart models.Cart
	if err := ctx.ShouldBindJSON(&cart); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdCart, err := c.cartService.AddToCart(cart)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, createdCart)
}

// PUT /cart/:id
func (c *CartController) UpdateCart(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var cart models.Cart
	if err := ctx.ShouldBindJSON(&cart); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cart.ID = uint(id)
	updatedCart, err := c.cartService.UpdateCart(cart)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedCart)
}

// DELETE /cart/:id
func (c *CartController) DeleteCart(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := c.cartService.RemoveFromCart(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "item removed from cart"})
}
