package system

import (
	"github.com/gjssss/soybean-admin-go/global"
	"github.com/gjssss/soybean-admin-go/models/system"
)

type MenuRepository struct{}

func (c *MenuRepository) GetMenus() ([]system.Menu, error) {
	var menus []system.Menu
	err := global.DB.Find(&menus).Error
	return menus, err
}

func (c *MenuRepository) GetMenusByUserId(userId uint) ([]system.Menu, error) {
	var menus []system.Menu
	err := global.DB.Raw(`
        SELECT m.* FROM menus m
        JOIN role_menus rm ON m.id = rm.menu_id
        JOIN user_roles ur ON rm.role_id = ur.role_id
        WHERE ur.user_id = ?
    `, userId).Scan(&menus).Error
	return menus, err
}

func (c *MenuRepository) GetMenusByRoleId(roleId uint) ([]system.Menu, error) {
	var menus []system.Menu
	err := global.DB.Raw(`
		SELECT m.* FROM menus m
		JOIN role_menus rm ON m.id = rm.menu_id
		WHERE rm.role_id = ?
	`, roleId).Scan(&menus).Error
	return menus, err
}

func (c *MenuRepository) CreateMenu(menu *system.Menu) error {
	return global.DB.Create(menu).Error
}

func (c *MenuRepository) MenuExists(menuId uint) (bool, error) {
	var count int64
	err := global.DB.Model(&system.Menu{}).Where("id = ?", menuId).Count(&count).Error
	return count > 0, err
}

func (c *MenuRepository) UpdateMenu(menu *system.Menu) error {
	return global.DB.Model(menu).Updates(menu).Error
}

func (c *MenuRepository) DeleteMenu(menuId uint) error {
	return global.DB.Delete(&system.Menu{}, menuId).Error
}

func (c *MenuRepository) HasChildren(menuId uint) (bool, error) {
	var count int64
	err := global.DB.Model(&system.Menu{}).Where("parent_id = ?", menuId).Count(&count).Error
	return count > 0, err
}

func (c *MenuRepository) GetMenuById(menuId uint) (*system.Menu, error) {
	var menu system.Menu
	err := global.DB.Where("id = ?", menuId).First(&menu).Error
	return &menu, err
}
