package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/global"
	"github.com/gjssss/soybean-admin-go/models"
	"github.com/gjssss/soybean-admin-go/routes/system"
)

func Init(r *gin.Engine) {
	system.Init(r)
	r.GET("/init", func(c *gin.Context) {
		models.InitDatabase(global.DB)
		c.JSON(200, gin.H{
			"message": "init success",
		})
	})
}
