package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/controllers"
	"github.com/gjssss/soybean-admin-go/middlewares"
)

var (
	UploadControllers = controllers.System.Upload
)

// UploadRoutes 上传相关路由
func UploadRoutes(r *gin.Engine) {
	uploadGroup := r.Group("/upload")
	{
		// 需要JWT认证
		uploadGroup.Use(middlewares.AuthMiddleware())

		// 获取上传凭证
		uploadGroup.GET("/token/aws", UploadControllers.GetUploadToken)
	}
}
