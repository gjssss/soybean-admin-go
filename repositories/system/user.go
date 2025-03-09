package system

import (
	"github.com/gjssss/soybean-admin-go/global"
	"github.com/gjssss/soybean-admin-go/models/system"
)

type UserRepository struct{}

func (r *UserRepository) FindAll() ([]system.User, error) {
	var users []system.User
	if err := global.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) Create(user system.User) (system.User, error) {
	if err := global.DB.Create(&user).Error; err != nil {
		return system.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Update(user system.User) (system.User, error) {
	if err := global.DB.Save(&user).Error; err != nil {
		return system.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Delete(user system.User) error {
	if err := global.DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindById(id uint) (system.User, error) {
	var user system.User
	if err := global.DB.Preload("Roles").Preload("Buttons").First(&user, id).Error; err != nil {
		return system.User{}, err
	}
	return user, nil
}

func (r *UserRepository) FindByUsername(username string) (system.User, error) {
	var user system.User
	if err := global.DB.Where("user_name = ?", username).First(&user).Error; err != nil {
		return system.User{}, err
	}
	return user, nil
}
