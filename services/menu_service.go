package services

import (
	"food-ordering/models"
	"food-ordering/repositories"
)

type MenuService interface {
	GetAllMenus() ([]models.Menu, error)
	GetMenuByID(id uint) (models.Menu, error)
	CreateMenu(menu models.Menu) (models.Menu, error)
}

type menuService struct {
	menuRepo repositories.MenuRepository
}

func NewMenuService(menuRepo repositories.MenuRepository) MenuService {
	return &menuService{menuRepo}
}

func (s *menuService) GetAllMenus() ([]models.Menu, error) {
	return s.menuRepo.FindAll()
}

func (s *menuService) GetMenuByID(id uint) (models.Menu, error) {
	return s.menuRepo.FindByID(id)
}

func (s *menuService) CreateMenu(menu models.Menu) (models.Menu, error) {
	return s.menuRepo.Create(menu)
}
