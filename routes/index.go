package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/routes/system"
)

func Init(r *gin.Engine) {
	system.Init(r)
}
