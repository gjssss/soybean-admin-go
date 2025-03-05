package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/controllers"
)

var (
	UserControllers = controllers.System.User
)

func Init(r *gin.Engine) {
	UserRoutes(r)
}
