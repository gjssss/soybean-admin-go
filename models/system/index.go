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

		Api{},
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
			Constant:  true,
		},
		// 系统管理及其子菜单
		{
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
			Status:    "1",
			ParentID:  2,
			MenuType:  "2",
			MenuName:  "按钮管理",
			RouteName: "manage_button",
			RoutePath: "/manage/button",
			Component: "view.manage_button",
			Order:     5,
			I18nKey:   "route.manage_button",
			Icon:      "carbon:button-centered",
			IconType:  "1",
		},
		{
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
		{
			Status:     "1",
			ParentID:   0,
			MenuType:   "2",
			MenuName:   "403",
			RouteName:  "403",
			RoutePath:  "/403",
			Component:  "layout.blank$view.403",
			Order:      6,
			I18nKey:    "route.403",
			Constant:   true,
			HideInMenu: true,
		},
		{
			Status:     "1",
			ParentID:   0,
			MenuType:   "2",
			MenuName:   "404",
			RouteName:  "404",
			RoutePath:  "/404",
			Component:  "layout.blank$view.404",
			Order:      7,
			I18nKey:    "route.404",
			Constant:   true,
			HideInMenu: true,
		},
		{
			Status:     "1",
			ParentID:   0,
			MenuType:   "2",
			MenuName:   "500",
			RouteName:  "500",
			RoutePath:  "/500",
			Component:  "layout.blank$view.500",
			Order:      8,
			I18nKey:    "route.500",
			Constant:   true,
			HideInMenu: true,
		},
		{
			Status:     "1",
			ParentID:   0,
			MenuType:   "2",
			MenuName:   "iframe-page",
			RouteName:  "iframe-page",
			RoutePath:  "/iframe-page/:url",
			Component:  "layout.base$view.iframe-page",
			Order:      9,
			I18nKey:    "route.iframe-page",
			Constant:   true,
			HideInMenu: true,
			KeepAlive:  true,
		},
		{
			Status:     "1",
			ParentID:   0,
			MenuType:   "2",
			MenuName:   "login",
			RouteName:  "login",
			RoutePath:  "/login/:module(pwd-login|code-login|register|reset-pwd|bind-wechat)?",
			Component:  "layout.blank$view.login",
			Order:      10,
			I18nKey:    "route.login",
			Constant:   true,
			HideInMenu: true,
		},
	}

	// API
	apis := []Api{
		{Path: "/auth/login", Method: "POST", Group: "权限接口"},
		{Path: "/auth/getUserInfo", Method: "GET", Group: "权限接口"},
		{Path: "/auth/refreshToken", Method: "POST", Group: "权限接口"},

		{Path: "/users", Method: "GET", Group: "用户管理"},
		{Path: "/users", Method: "POST", Group: "用户管理"},
		{Path: "/users/password", Method: "POST", Group: "用户管理"},
		{Path: "/users/delete", Method: "POST", Group: "用户管理"},
		{Path: "/users/batchDelete", Method: "POST", Group: "用户管理"},
		{Path: "/users/roles", Method: "POST", Group: "用户管理"},
		{Path: "/users/checkUsername", Method: "GET", Group: "用户管理"},
		{Path: "/users/roles", Method: "GET", Group: "用户管理"},

		{Path: "/roles/all", Method: "GET", Group: "角色管理"},
		{Path: "/roles", Method: "GET", Group: "角色管理"},
		{Path: "/roles", Method: "POST", Group: "角色管理"},
		{Path: "/roles/update", Method: "POST", Group: "角色管理"},
		{Path: "/roles/delete", Method: "POST", Group: "角色管理"},
		{Path: "/roles/batchDelete", Method: "POST", Group: "角色管理"},
		{Path: "/roles/menus", Method: "POST", Group: "角色管理"},
		{Path: "/roles/buttons", Method: "POST", Group: "角色管理"},

		{Path: "/menus", Method: "GET", Group: "菜单管理"},
		{Path: "/menus/user", Method: "GET", Group: "菜单管理"},
		{Path: "/menus/role", Method: "GET", Group: "菜单管理"},
		{Path: "/menus", Method: "POST", Group: "菜单管理"},
		{Path: "/menus/update", Method: "POST", Group: "菜单管理"},
		{Path: "/menus/delete", Method: "POST", Group: "菜单管理"},
		{Path: "/menus/batchDelete", Method: "POST", Group: "菜单管理"},
		{Path: "/menus/constant", Method: "GET", Group: "菜单管理"},

		{Path: "/buttons", Method: "GET", Group: "按钮管理"},
		{Path: "/buttons/role", Method: "GET", Group: "按钮管理"},
		{Path: "/buttons/user", Method: "GET", Group: "按钮管理"},
		{Path: "/buttons", Method: "POST", Group: "按钮管理"},
		{Path: "/buttons/update", Method: "POST", Group: "按钮管理"},
		{Path: "/buttons/delete", Method: "POST", Group: "按钮管理"},
		{Path: "/buttons/batchDelete", Method: "POST", Group: "按钮管理"},

		{Path: "/apis", Method: "GET", Group: "接口管理"},
		{Path: "/apis/role", Method: "GET", Group: "接口管理"},
		{Path: "/apis", Method: "POST", Group: "接口管理"},
		{Path: "/apis/update", Method: "POST", Group: "接口管理"},
		{Path: "/apis/delete", Method: "POST", Group: "接口管理"},
		{Path: "/apis/role", Method: "POST", Group: "接口管理"},
	}

	// Role
	role1 := Role{RoleName: "R_SUPER", RoleDesc: "超级管理员", Menus: menus, Buttons: buttons, Apis: apis}
	role2 := Role{RoleName: "R_USER", RoleDesc: "用户", Menus: make([]Menu, 0), Buttons: make([]Button, 0), Apis: make([]Api, 0)}

	// User
	users := []*User{
		{UserName: "admin", Password: encodePassword("123123123"), Roles: []Role{role1}},
		{UserName: "user", Password: encodePassword("123123123"), Roles: []Role{role2}},
	}

	db.Create(users)
}

func encodePassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}
