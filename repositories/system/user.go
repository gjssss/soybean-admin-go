package system

import (
	"github.com/gjssss/soybean-admin-go/global"
	"github.com/gjssss/soybean-admin-go/models"
)

type UserRepository struct{}

func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	if err := global.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) Create(user models.User) (models.User, error) {
	if err := global.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
