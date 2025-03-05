package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/controllers"
)

func UserRoutes(r *gin.Engine, controller *controllers.UserController) {
	users := r.Group("/users")
	{
		users.GET("", controller.GetAllUsers)
		users.POST("", controller.CreateUser)
	}
}
