package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/middlewares"
)

func MenuRoutes(r *gin.Engine) {
	menus := r.Group("/menus")
	menus.Use(middlewares.AuthMiddleware())
	{
		menus.GET("", SystemControllers.Menu.GetMenus)
		menus.GET("/user", SystemControllers.Menu.GetUserMenus)
		menus.GET("/role", SystemControllers.Menu.GetMenusByRoleId)
		menus.GET("/constant", SystemControllers.Menu.GetConstantMenus)
		menus.POST("", SystemControllers.Menu.CreateMenu)
		menus.POST("/update", SystemControllers.Menu.UpdateMenu)
		menus.POST("/delete", SystemControllers.Menu.DeleteMenu)
		menus.POST("/batchDelete", SystemControllers.Menu.BatchDeleteMenu)
	}
}
