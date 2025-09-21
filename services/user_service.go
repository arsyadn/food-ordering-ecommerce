package services

import (
	"errors"
	"food-ordering/models"
	"food-ordering/repositories"
	"food-ordering/utils"

	"gorm.io/gorm"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		Repo: repositories.NewUserRepository(db),
	}
}

func (s *UserService) Register(user *models.User) (string, uint, error) {
	// hash password
	if err := user.HashPassword(user.Password); err != nil {
		return "", 0, errors.New("failed to hash password")
	}

	// save user
	if err := s.Repo.Create(user); err != nil {
		return "", 0, errors.New("failed to create user")
	}

	// generate token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", 0, errors.New("failed to generate token")
	}

	return token, user.ID, nil
}

func (s *UserService) Login(req *models.LoginRequest) (string, uint, error) {
	user, err := s.Repo.FindByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", 0, errors.New("invalid email or password")
		}
		return "", 0, err
	}

	if err := user.CheckPassword(req.Password); err != nil {
		return "", 0, errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", 0, errors.New("failed to generate token")
	}

	return token, user.ID, nil
}
