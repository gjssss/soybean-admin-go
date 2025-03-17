package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/middlewares"
)

func ButtonRoutes(r *gin.Engine) {
	buttons := r.Group("/buttons")
	buttons.Use(middlewares.AuthMiddleware())
	buttons.Use(middlewares.ApiMiddleware())
	{
		buttons.GET("", SystemControllers.Button.GetButtons)
		buttons.GET("/role", SystemControllers.Button.GetButtonsByRoleId)
		buttons.GET("/user", SystemControllers.Button.GetUserButtons)
		buttons.POST("", SystemControllers.Button.CreateButton)
		buttons.POST("/update", SystemControllers.Button.UpdateButton)
		buttons.POST("/delete", SystemControllers.Button.DeleteButton)
		buttons.POST("/batchDelete", SystemControllers.Button.BatchDeleteButton)
	}
}
