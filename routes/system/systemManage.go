package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/middlewares"
)

func SystemManageRoutes(r *gin.Engine) {
	routes := r.Group("/systemManage")
	routes.Use(middlewares.AuthMiddleware())
	{
		routes.GET("/getMenuList", SystemControllers.Menu.GetMenus)
		routes.GET("/getRoleList", SystemControllers.Role.GetRoles)
		routes.GET("/getUserList", SystemControllers.User.GetAllUsers)
		routes.GET("/getAllRoles", SystemControllers.Role.GetAllRoles)
		routes.GET("/getRoleMenu", SystemControllers.Menu.GetMenusByRoleId)
	}
}
