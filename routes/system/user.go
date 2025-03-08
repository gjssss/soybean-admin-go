package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/middlewares"
)

func UserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("", middlewares.AuthMiddleware(), SystemControllers.User.GetAllUsers)
		users.POST("", SystemControllers.User.CreateUser)
	}
}
