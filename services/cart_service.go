package services

import (
	"food-ordering/models"
	"food-ordering/repositories"
)

type CartService interface {
	GetUserCart(userID uint) ([]models.Cart, error)
	AddToCart(cart models.Cart) (models.Cart, error)
	UpdateCart(cart models.Cart) (models.Cart, error)
	RemoveFromCart(id uint) error
}

type cartService struct {
	cartRepo repositories.CartRepository
}

func NewCartService(cartRepo repositories.CartRepository) CartService {
	return &cartService{cartRepo}
}

func (s *cartService) GetUserCart(userID uint) ([]models.Cart, error) {
	return s.cartRepo.FindAll(userID)
}

func (s *cartService) AddToCart(cart models.Cart) (models.Cart, error) {
	return s.cartRepo.Create(cart)
}

func (s *cartService) UpdateCart(cart models.Cart) (models.Cart, error) {
	return s.cartRepo.Update(cart)
}

func (s *cartService) RemoveFromCart(id uint) error {
	return s.cartRepo.Delete(id)
}
