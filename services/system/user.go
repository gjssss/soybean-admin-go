package system

import (
	"github.com/gjssss/soybean-admin-go/models"
)

type UserService struct {
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return SystemRepositories.User.FindAll()
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	return SystemRepositories.User.Create(user)
}
