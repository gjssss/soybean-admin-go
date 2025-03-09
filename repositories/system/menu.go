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
