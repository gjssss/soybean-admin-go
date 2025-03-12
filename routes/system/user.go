package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/middlewares"
)

func UserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	users.Use(middlewares.AuthMiddleware())
	{
		users.GET("", SystemControllers.User.GetAllUsers)
		users.POST("", SystemControllers.User.CreateUser)
		users.POST("/password", SystemControllers.User.UpdateUserPassword)
		users.POST("/delete", SystemControllers.User.DeleteUser)
		users.POST("/batchDelete", SystemControllers.User.BatchDeleteUser)
		users.POST("/roles", SystemControllers.User.UpdateUserRoles)
		users.GET("/checkUsername", SystemControllers.User.CheckUserNameExists)
		users.GET("/roles", SystemControllers.User.GetUserRoles)
	}
}
