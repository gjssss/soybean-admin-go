package system

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("", UserControllers.GetAllUsers)
		users.POST("", UserControllers.CreateUser)
	}
}
