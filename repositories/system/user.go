package system

import (
	"github.com/gjssss/soybean-admin-go/global"
	"github.com/gjssss/soybean-admin-go/models/system"
	"github.com/gjssss/soybean-admin-go/utils"
)

type UserRepository struct{}

func (r *UserRepository) FindAll(page utils.PaginationParam) ([]system.User, int64, error) {
	var users []system.User
	if err := global.DB.Offset(page.Size * (page.Current - 1)).Limit(page.Size).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	var total int64
	if err := global.DB.Model(&system.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil
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

func (r *UserRepository) UpdatePassword(user system.User) error {
	if err := global.DB.Model(&user).Update("password", user.Password).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Delete(user system.User) error {
	println("delete", user.ID)
	if err := global.DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindById(id uint) (system.User, error) {
	var user system.User
	if err := global.DB.Preload("Roles").First(&user, id).Error; err != nil {
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

func (r *UserRepository) UpdateUserRoles(userID uint, roleIDs []uint) error {
	// 先清除用户现有的角色关联
	if err := global.DB.Model(&system.User{ID: userID}).Association("Roles").Clear(); err != nil {
		return err
	}

	// 添加新的角色关联
	var roles []system.Role
	for _, roleID := range roleIDs {
		roles = append(roles, system.Role{ID: roleID})
	}

	return global.DB.Model(&system.User{ID: userID}).Association("Roles").Append(roles)
}
