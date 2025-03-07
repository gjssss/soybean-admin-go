package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/middlewares"
)

func AuthRoutes(r *gin.Engine) {
	{
		auth := r.Group("/auth")
		auth.POST("/login", UserControllers.Login)
	}
	{
		auth := r.Group("/auth")
		auth.Use(middlewares.AuthMiddleware())
		auth.GET("/getUserInfo", UserControllers.GetUserInfo)
		auth.POST("/refreshToken", UserControllers.RefreshToken)
	}
}
