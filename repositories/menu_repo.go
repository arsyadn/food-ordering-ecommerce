package repositories

import (
	"food-ordering/models"

	"gorm.io/gorm"
)

type MenuRepository interface {
	FindAll() ([]models.Menu, error)
	FindByID(id uint) (models.Menu, error)
	Create(menu models.Menu) (models.Menu, error)
	Update(menu models.Menu) (models.Menu, error)
	Delete(id uint) error
}

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepository{db: db}
}

func (r *menuRepository) FindAll() ([]models.Menu, error) {
	var menus []models.Menu
	err := r.db.Find(&menus).Error
	return menus, err
}

func (r *menuRepository) FindByID(id uint) (models.Menu, error) {
	var menu models.Menu
	err := r.db.First(&menu, id).Error
	return menu, err
}

func (r *menuRepository) Create(menu models.Menu) (models.Menu, error) {
	err := r.db.Create(&menu).Error
	return menu, err
}

func (r *menuRepository) Update(menu models.Menu) (models.Menu, error) {
	err := r.db.Save(&menu).Error
	return menu, err
}

func (r *menuRepository) Delete(id uint) error {
	return r.db.Delete(&models.Menu{}, id).Error
}
