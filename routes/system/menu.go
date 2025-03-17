package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/middlewares"
)

func MenuRoutes(r *gin.Engine) {
	{
		menus := r.Group("/menus")
		menus.Use(middlewares.AuthMiddleware())
		menus.GET("/user", SystemControllers.Menu.GetUserMenus)
		{
			admin_menus := menus.Group("")
			admin_menus.Use(middlewares.ApiMiddleware())
			menus.GET("", SystemControllers.Menu.GetMenus)
			menus.GET("/role", SystemControllers.Menu.GetMenusByRoleId)
			menus.POST("", SystemControllers.Menu.CreateMenu)
			menus.POST("/update", SystemControllers.Menu.UpdateMenu)
			menus.POST("/delete", SystemControllers.Menu.DeleteMenu)
			menus.POST("/batchDelete", SystemControllers.Menu.BatchDeleteMenu)
		}
	}
	{
		menus := r.Group("/menus")
		menus.GET("/constant", SystemControllers.Menu.GetConstantMenus)
	}
}
