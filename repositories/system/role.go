package system

import (
	"github.com/gjssss/soybean-admin-go/global"
	"github.com/gjssss/soybean-admin-go/models/system"
	"github.com/gjssss/soybean-admin-go/utils"
)

type RoleRepository struct{}

func (c *RoleRepository) GetRoles(page utils.PaginationParam) ([]system.Role, int64, error) {
	var roles []system.Role
	var count int64
	err := global.DB.Offset(page.Size * (page.Current - 1)).Limit(page.Size).Find(&roles).Error
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
	tx := global.DB.Begin()

	// Delete associations in user_roles table first
	if err := tx.Exec("DELETE FROM user_roles WHERE role_id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete associations in role_menus table
	if err := tx.Exec("DELETE FROM role_menus WHERE role_id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete associations in role_buttons table
	if err := tx.Exec("DELETE FROM role_buttons WHERE role_id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete associations in role_api table (if it exists)
	if err := tx.Exec("DELETE FROM role_api WHERE role_id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Then delete the role itself
	if err := tx.Delete(&system.Role{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (c *RoleRepository) BatchDeleteRole(ids []uint) error {
	tx := global.DB.Begin()

	// Delete associations in user_roles table first
	if err := tx.Exec("DELETE FROM user_roles WHERE role_id IN ?", ids).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete associations in role_menus table
	if err := tx.Exec("DELETE FROM role_menus WHERE role_id IN ?", ids).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete associations in role_buttons table
	if err := tx.Exec("DELETE FROM role_buttons WHERE role_id IN ?", ids).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete associations in role_api table (if it exists)
	if err := tx.Exec("DELETE FROM role_api WHERE role_id IN ?", ids).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Then delete the roles themselves
	if err := tx.Delete(&system.Role{}, ids).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
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

func (r *RoleRepository) RoleExists(roleID uint) (bool, error) {
	var count int64
	err := global.DB.Model(&system.Role{}).Where("id = ?", roleID).Count(&count).Error
	return count > 0, err
}
