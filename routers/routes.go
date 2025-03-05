package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/controllers"
)

func SetupRoutes(r *gin.Engine, userController *controllers.UserController) {
	UserRoutes(r, userController)
}
