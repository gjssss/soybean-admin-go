package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/middlewares"
)

func ButtonRoutes(r *gin.Engine) {
	buttons := r.Group("/buttons")
	buttons.Use(middlewares.AuthMiddleware())
	{
		buttons.GET("", SystemControllers.Button.GetButtons)
		buttons.GET("/role", SystemControllers.Button.GetButtonsByRoleId)
		buttons.GET("/user", SystemControllers.Button.GetUserButtons)
	}
}
