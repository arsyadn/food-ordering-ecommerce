package services

import (
	"errors"
	"food-ordering/models"
	"food-ordering/repositories"

	"gorm.io/gorm"
)

type OrderService interface {
	Checkout(userID uint) (models.Order, error)
	GetUserOrders(userID uint) ([]models.Order, error)
}

type orderService struct {
	orderRepo repositories.OrderRepository
	cartRepo  repositories.CartRepository
	menuRepo  repositories.MenuRepository
	db        *gorm.DB
}

func NewOrderService(orderRepo repositories.OrderRepository, cartRepo repositories.CartRepository, menuRepo repositories.MenuRepository, db *gorm.DB) OrderService {
	return &orderService{orderRepo, cartRepo, menuRepo, db}
}

func (s *orderService) Checkout(userID uint) (models.Order, error) {
	var order models.Order

	// ambil isi cart user
	carts, err := s.cartRepo.FindAll(userID)
	if err != nil {
		return order, err
	}
	if len(carts) == 0 {
		return order, errors.New("cart is empty")
	}

	// jalankan transaksi
	err = s.db.Transaction(func(tx *gorm.DB) error {
		var total float64
		var orderItems []models.OrderItem

		for _, cart := range carts {
			menu, err := s.menuRepo.FindByID(cart.MenuID)
			if err != nil {
				return err
			}
			if menu.Stock < cart.Quantity {
				return errors.New("insufficient stock for " + menu.Name)
			}

			// buat order item
			subtotal := float64(cart.Quantity) * menu.Price
			orderItems = append(orderItems, models.OrderItem{
				MenuID:   menu.ID,
				Quantity: cart.Quantity,
				Price:    menu.Price,
				Subtotal: subtotal,
			})
			total += subtotal

			// update stok
			menu.Stock -= cart.Quantity
			if _, err := s.menuRepo.Update(menu); err != nil {
				return err
			}
		}

		// buat order
		order = models.Order{
			UserID: userID,
			Total:  total,
			Items:  orderItems,
		}
		if _, err := s.orderRepo.Create(order); err != nil {
			return err
		}

		// kosongkan cart
		for _, cart := range carts {
			if err := s.cartRepo.Delete(cart.ID); err != nil {
				return err
			}
		}

		return nil
	})

	return order, err
}

func (s *orderService) GetUserOrders(userID uint) ([]models.Order, error) {
	return s.orderRepo.FindByUserID(userID)
}
