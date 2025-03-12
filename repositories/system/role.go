package system

import (
	"github.com/gjssss/soybean-admin-go/global"
	"github.com/gjssss/soybean-admin-go/models/common"
	"github.com/gjssss/soybean-admin-go/models/system"
)

type RoleRepository struct{}

func (c *RoleRepository) GetRoles(page common.PaginationParam) ([]system.Role, int64, error) {
	var roles []system.Role
	var count int64
	err := global.DB.Offset(page.PageSize * (page.Current - 1)).Limit(page.PageSize).Find(&roles).Error
	if err != nil {
		return nil, count, err
	}
	err = global.DB.Model(&system.Role{}).Count(&count).Error
	if err != nil {
		return nil, count, err
	}
	return roles, count, err
}

func (c *RoleRepository) GetAllRoles() ([]system.Role, error) {
	var roles []system.Role
	err := global.DB.Find(&roles).Error
	return roles, err
}

func (c *RoleRepository) CreateRole(role *system.Role) error {
	return global.DB.Create(role).Error
}

func (c *RoleRepository) UpdateRole(role *system.Role) error {
	return global.DB.Save(role).Error
}

func (c *RoleRepository) DeleteRole(id uint) error {
	return global.DB.Delete(&system.Role{}, id).Error
}

func (c *RoleRepository) BatchDeleteRole(ids []uint) error {
	return global.DB.Delete(&system.Role{}, ids).Error
}

func (c *RoleRepository) UpdateRoleMenu(roleID uint, menuIDs []uint) error {
	// 先清除角色的所有菜单
	role := system.Role{ID: roleID}
	if err := global.DB.Model(&role).Association("Menus").Clear(); err != nil {
		return err
	}

	// 如果没有新菜单要添加，则直接返回
	if len(menuIDs) == 0 {
		return nil
	}

	// 添加新菜单
	var menus []system.Menu
	if err := global.DB.Find(&menus, menuIDs).Error; err != nil {
		return err
	}
	return global.DB.Model(&role).Association("Menus").Append(&menus)
}

func (c *RoleRepository) UpdateRoleButton(roleID uint, buttonIDs []uint) error {
	// 先清除角色的所有按钮权限
	role := system.Role{ID: roleID}
	if err := global.DB.Model(&role).Association("Buttons").Clear(); err != nil {
		return err
	}

	// 如果没有新按钮要添加，则直接返回
	if len(buttonIDs) == 0 {
		return nil
	}

	// 添加新按钮权限
	var buttons []system.Button
	if err := global.DB.Find(&buttons, buttonIDs).Error; err != nil {
		return err
	}
	return global.DB.Model(&role).Association("Buttons").Append(&buttons)
}
