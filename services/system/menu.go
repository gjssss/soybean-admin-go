package system

import (
	"errors"

	"github.com/gjssss/soybean-admin-go/models/system"
)

type MenuService struct{}

func (s *MenuService) GetMenusByUserId(userId uint) ([]*system.Menu, error) {
	dbMenu, err := SystemRepositories.Menu.GetMenusByUserId(userId)
	if err != nil {
		return make([]*system.Menu, 0), err
	}
	var menus []*system.Menu
	var cache = make(map[uint]*system.Menu)
	for _, menu := range dbMenu {
		menu.Children = make([]system.Menu, 0)
		if menu.ParentID == 0 {
			menus = append(menus, &menu)
			cache[menu.ID] = &menu
		} else {
			if parent, ok := cache[menu.ParentID]; ok {
				parent.Children = append(parent.Children, menu)
			} else {
				err = errors.New("菜单格式错误，父菜单不存在")
			}
		}
	}
	return menus, err
}
