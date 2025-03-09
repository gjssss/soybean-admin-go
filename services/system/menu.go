package system

import (
	"errors"

	"github.com/gjssss/soybean-admin-go/models/system"
)

type MenuService struct{}

func nestedMenu(menus *[]system.Menu) ([]*system.Menu, error) {
	var _menus []*system.Menu
	var err error
	var cache = make(map[uint]*system.Menu)
	for _, menu := range *menus {
		menu.Children = make([]system.Menu, 0)
		if menu.ParentID == 0 {
			_menus = append(_menus, &menu)
			cache[menu.ID] = &menu
		} else {
			if parent, ok := cache[menu.ParentID]; ok {
				parent.Children = append(parent.Children, menu)
			} else {
				err = errors.New("菜单格式错误，父菜单不存在")
			}
		}
	}
	return _menus, err
}

func (s *MenuService) GetMenusByUserId(userId uint) ([]*system.Menu, error) {
	dbMenu, err := SystemRepositories.Menu.GetMenusByUserId(userId)
	if err != nil {
		return make([]*system.Menu, 0), err
	}
	menus, err := nestedMenu(&dbMenu)
	return menus, err
}

func (s *MenuService) GetMenus() ([]*system.Menu, error) {
	dbMenu, err := SystemRepositories.Menu.GetMenus()
	if err != nil {
		return make([]*system.Menu, 0), err
	}
	menus, err := nestedMenu(&dbMenu)
	return menus, err
}
