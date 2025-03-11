package system

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		User{},

		Button{},

		Menu{},
		MenuQuery{},

		Role{},
	)
}

func InitDatabase(db *gorm.DB) {
	AutoMigrate(db)
	// m := db.Migrator()

	// Button
	buttons := []Button{
		{Code: "B_CODE1", Desc: "按钮1"},
		{Code: "B_CODE2", Desc: "按钮2"},
		{Code: "B_CODE3", Desc: "按钮3"},
	}

	// Menu
	menus := []Menu{
		{
			ID:        1,
			Status:    "1",
			ParentID:  0,
			MenuType:  "2",
			MenuName:  "首页",
			RouteName: "home",
			RoutePath: "/home",
			Component: "layout.base$view.home",
			Order:     1,
			I18nKey:   "route.home",
			Icon:      "mdi:monitor-dashboard",
			IconType:  "1",
		},
		// 系统管理及其子菜单
		{
			ID:        2,
			Status:    "1",
			ParentID:  0,
			MenuType:  "1",
			MenuName:  "系统管理",
			RouteName: "manage",
			RoutePath: "/manage",
			Component: "layout.base",
			Order:     5,
			I18nKey:   "route.manage",
			Icon:      "carbon:cloud-service-management",
			IconType:  "1",
		},
		{
			ID:        3,
			Status:    "1",
			ParentID:  2,
			MenuType:  "2",
			MenuName:  "用户管理",
			RouteName: "manage_user",
			RoutePath: "/manage/user",
			Component: "view.manage_user",
			Order:     1,
			I18nKey:   "route.manage_user",
			Icon:      "ic:round-manage-accounts",
			IconType:  "1",
		},
		{
			ID:        4,
			Status:    "1",
			ParentID:  2,
			MenuType:  "2",
			MenuName:  "角色管理",
			RouteName: "manage_role",
			RoutePath: "/manage/role",
			Component: "view.manage_role",
			Order:     2,
			I18nKey:   "route.manage_role",
			Icon:      "carbon:user-role",
			IconType:  "1",
		},
		{
			ID:        5,
			Status:    "1",
			ParentID:  2,
			MenuType:  "2",
			MenuName:  "菜单管理",
			RouteName: "manage_menu",
			RoutePath: "/manage/menu",
			Component: "view.manage_menu",
			Order:     3,
			I18nKey:   "route.manage_menu",
			Icon:      "material-symbols:route",
			IconType:  "1",
		},
		{
			ID:         6,
			Status:     "1",
			ParentID:   2,
			MenuType:   "2",
			MenuName:   "用户详情",
			RouteName:  "manage_user-detail",
			RoutePath:  "/manage/user-detail/:id",
			Component:  "view.manage_user-detail",
			Order:      4,
			I18nKey:    "route.manage_user-detail",
			HideInMenu: true,
		},
	}

	// Role
	roles := []Role{
		{RoleName: "R_SUPER", RoleDesc: "超级管理员", Menus: menus, Buttons: buttons},
	}

	// User
	users := []*User{
		{UserName: "admin", Password: encodePassword("123123123"), Roles: roles},
	}
	db.Create(users)
}

func encodePassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}
