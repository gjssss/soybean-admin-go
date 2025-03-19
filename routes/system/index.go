package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/controllers"
)

var (
	SystemControllers = controllers.System
)

func Init(r *gin.Engine) {
	AuthRoutes(r)
	UserRoutes(r)
	RoleRoutes(r)
	MenuRoutes(r)
	ButtonRoutes(r)
	ApiRoutes(r)
	UploadRoutes(r)
}
