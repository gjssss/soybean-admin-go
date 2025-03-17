package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/middlewares"
)

func ApiRoutes(r *gin.Engine) {
	apis := r.Group("/apis")
	apis.Use(middlewares.AuthMiddleware())
	{
		apis.GET("", SystemControllers.Api.GetAllApis)
		apis.GET("/role", SystemControllers.Api.GetApisByRoleID)
		apis.POST("", SystemControllers.Api.CreateApi)
		apis.POST("/update", SystemControllers.Api.UpdateApi)
		apis.POST("/delete", SystemControllers.Api.DeleteApi)
		apis.POST("/role", SystemControllers.Api.UpdateRoleApi)
	}
}
