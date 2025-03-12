package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/middlewares"
)

func RoleRoutes(r *gin.Engine) {
	roles := r.Group("/roles")
	roles.Use(middlewares.AuthMiddleware())
	{
		roles.GET("/all", SystemControllers.Role.GetAllRoles)
		roles.GET("", SystemControllers.Role.GetRoles)
		roles.POST("", SystemControllers.Role.CreateRole)
		roles.POST("/update", SystemControllers.Role.UpdateRole)
		roles.POST("/delete", SystemControllers.Role.DeleteRole)
		roles.POST("/batchDelete", SystemControllers.Role.BatchDeleteRole)
		roles.POST("/menus", SystemControllers.Role.UpdateRoleMenu)
		roles.POST("/buttons", SystemControllers.Role.UpdateRoleButton)
	}
}
