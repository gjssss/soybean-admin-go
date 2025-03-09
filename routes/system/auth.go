package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/middlewares"
)

func AuthRoutes(r *gin.Engine) {
	{
		auth := r.Group("/auth")
		auth.POST("/login", SystemControllers.User.Login)
	}
	{
		auth := r.Group("/auth")
		auth.Use(middlewares.AuthMiddleware())
		auth.GET("/getUserInfo", SystemControllers.User.GetUserInfo)
		auth.POST("/refreshToken", SystemControllers.User.RefreshToken)
		auth.GET("/menu", SystemControllers.Menu.GetUserMenus)
	}
}
