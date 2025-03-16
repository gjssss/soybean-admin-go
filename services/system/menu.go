package system

import (
	"errors"
	"fmt"

	"github.com/gjssss/soybean-admin-go/models/system"
)

type MenuService struct{}

func nestedMenu(menus *[]system.Menu) ([]*system.Menu, error) {
	var _menus []*system.Menu
	var err error
	var cache = make(map[uint]*system.Menu)
	for _, menu := range *menus {
		menu.Children = make([]system.Menu, 0)
		cache[menu.ID] = &menu
	}
	for _, menu := range cache {
		if menu.ParentID != 0 {
			if parent, ok := cache[menu.ParentID]; ok {
				parent.Children = append(parent.Children, *menu)
			} else {
				err = errors.New("父菜单不存在")
			}
		} else {
			_menus = append(_menus, menu)
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

func (s *MenuService) GetConstantMenu() ([]*system.Menu, error) {
	dbMenu, err := SystemRepositories.Menu.GetConstantMenu()
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

func (s *MenuService) GetMenusByRoleId(roleId uint) ([]system.Menu, error) {
	dbMenu, err := SystemRepositories.Menu.GetMenusByRoleId(roleId)
	return dbMenu, err
}

func (s *MenuService) CreateMenu(menu *system.Menu) error {
	// 如果父ID不为0，检查父菜单是否存在
	if menu.ParentID != 0 {
		exists, err := SystemRepositories.Menu.MenuExists(menu.ParentID)
		if err != nil {
			return err
		}
		if !exists {
			return errors.New("父菜单不存在")
		}
	}

	return SystemRepositories.Menu.CreateMenu(menu)
}

func (s *MenuService) UpdateMenu(menu *system.Menu) error {
	// 检查菜单是否存在
	oldMenu, err := SystemRepositories.Menu.GetMenuById(menu.ID)
	if err != nil {
		return errors.New("菜单不存在")
	}

	// 如果修改了父ID，且父ID不为0，检查父菜单是否存在
	if menu.ParentID != oldMenu.ParentID && menu.ParentID != 0 {
		exists, err := SystemRepositories.Menu.MenuExists(menu.ParentID)
		if err != nil {
			return err
		}
		if !exists {
			return errors.New("父菜单不存在")
		}
	}

	return SystemRepositories.Menu.UpdateMenu(menu)
}

func (s *MenuService) DeleteMenu(menuId uint) error {
	// 检查菜单是否存在
	exists, err := SystemRepositories.Menu.MenuExists(menuId)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("菜单不存在")
	}

	// 检查是否有子菜单
	hasChildren, err := SystemRepositories.Menu.HasChildren(menuId)
	if err != nil {
		return err
	}
	if hasChildren {
		return errors.New("该菜单存在子菜单，不能删除")
	}

	return SystemRepositories.Menu.DeleteMenu(menuId)
}

func (s *MenuService) BatchDeleteMenu(menuIds []uint) error {
	for _, id := range menuIds {
		err := s.DeleteMenu(id)
		if err != nil {
			return fmt.Errorf("删除菜单ID %d 时出错: %s", id, err.Error())
		}
	}
	return nil
}
