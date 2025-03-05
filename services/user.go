package services

import (
	"github.com/gjssss/soybean-admin-go/models"
	"github.com/gjssss/soybean-admin-go/repositories"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
}

type UserServiceImpl struct {
	userRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &UserServiceImpl{userRepo: repo}
}

func (s *UserServiceImpl) GetAllUsers() ([]models.User, error) {
	return s.userRepo.FindAll()
}

func (s *UserServiceImpl) CreateUser(user models.User) (models.User, error) {
	return s.userRepo.Create(user)
}
