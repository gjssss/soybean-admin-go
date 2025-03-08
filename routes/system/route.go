package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/middlewares"
)

func RouteRoutes(r *gin.Engine) {
	routes := r.Group("/route")
	routes.Use(middlewares.AuthMiddleware())
	{
		routes.GET("/getUserRoutes", SystemControllers.Menu.GetUserMenus)
	}
}
