package repositories

import (
	"food-ordering/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindAll(userID uint) ([]models.Cart, error)
	FindByID(id uint) (models.Cart, error)
	Create(cart models.Cart) (models.Cart, error)
	Update(cart models.Cart) (models.Cart, error)
	Delete(id uint) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) FindAll(userID uint) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Where("user_id = ?", userID).Find(&carts).Error
	return carts, err
}

func (r *cartRepository) FindByID(id uint) (models.Cart, error) {
	var cart models.Cart
	err := r.db.First(&cart, id).Error
	return cart, err
}

func (r *cartRepository) Create(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Error
	return cart, err
}

func (r *cartRepository) Update(cart models.Cart) (models.Cart, error) {
	err := r.db.Exec("UPDATE carts SET quantity = ?, updated_at = NOW() WHERE id = ?",  cart.Quantity, cart.ID).Error
	return cart, err
}

func (r *cartRepository) Delete(id uint) error {
	return r.db.Delete(&models.Cart{}, id).Error
}
