package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/middlewares"
)

func AuthRoutes(r *gin.Engine) {
	{
		// 无需认证的接口
		auth := r.Group("/auth")

		auth.POST("/login", SystemControllers.User.Login)
	}
	{
		// 需要认证的接口
		auth := r.Group("/auth")
		auth.Use(middlewares.AuthMiddleware())

		auth.GET("/getUserInfo", SystemControllers.User.GetUserInfo)
		auth.POST("/refreshToken", SystemControllers.User.RefreshToken)
	}
}
