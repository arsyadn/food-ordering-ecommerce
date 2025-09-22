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
		userID := ctx.GetUint("user_id")

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
	userID := ctx.GetUint("user_id")

	id, _ := strconv.Atoi(ctx.Param("id"))
	cart.MenuID = uint(id)
	cart.UserID = userID

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

func (c *CartController) GetUint(s string) any {
	panic("unimplemented")
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
