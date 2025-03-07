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

func (r *UserRepository) Update(user models.User) (models.User, error) {
	if err := global.DB.Save(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Delete(user models.User) error {
	if err := global.DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindById(id uint) (models.User, error) {
	var user models.User
	if err := global.DB.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) FindByUsername(username string) (models.User, error) {
	var user models.User
	if err := global.DB.Where("user_name = ?", username).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
