package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/middlewares"
)

func SystemManageRoutes(r *gin.Engine) {
	routes := r.Group("/systemManage")
	routes.Use(middlewares.AuthMiddleware())
	{
		routes.GET("/getRoleList", SystemControllers.Role.GetRoles)
		routes.POST("/createRole", SystemControllers.Role.CreateRole)
		routes.PUT("/updateRole", SystemControllers.Role.UpdateRole)
		routes.DELETE("/deleteRole", SystemControllers.Role.DeleteRole)
		routes.DELETE("/batchDeleteRole", SystemControllers.Role.BatchDeleteRole)

		routes.GET("/getAllRoles", SystemControllers.Role.GetAllRoles)
		routes.GET("/getRoleMenu", SystemControllers.Menu.GetMenusByRoleId)
		routes.PUT("/updateRoleMenu", SystemControllers.Role.UpdateRoleMenu)
		routes.GET("/getAllButton", SystemControllers.Button.GetButtons)
		routes.GET("/getRoleButton", SystemControllers.Button.GetButtonsByRoleId)
		routes.PUT("/updateRoleButton", SystemControllers.Role.UpdateRoleButton)

		routes.GET("/getMenuList", SystemControllers.Menu.GetMenus)
		routes.POST("/createMenu", SystemControllers.Menu.CreateMenu)
		routes.PUT("/updateMenu", SystemControllers.Menu.UpdateMenu)
		routes.DELETE("/deleteMenu", SystemControllers.Menu.DeleteMenu)
		routes.DELETE("/batchDeleteMenu", SystemControllers.Menu.BatchDeleteMenu)

		routes.GET("/getUserList", SystemControllers.User.GetAllUsers)
		routes.POST("/createUser", SystemControllers.User.CreateUser)
		routes.PUT("/updateUserPassword", SystemControllers.User.UpdateUserPassword)
		routes.PUT("/updateUserRoles", SystemControllers.User.UpdateUserRoles)
		routes.DELETE("/deleteUser", SystemControllers.User.DeleteUser)
		routes.DELETE("/batchDeleteUser", SystemControllers.User.BatchDeleteUser)

		routes.GET("/checkUserNameExists", SystemControllers.User.CheckUserNameExists)
		routes.GET("/getUserRoles", SystemControllers.User.GetUserRoles)
	}
}
