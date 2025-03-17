package main

import (
	"github.com/gjssss/soybean-admin-go/database"
	"github.com/gjssss/soybean-admin-go/global"
	"github.com/gjssss/soybean-admin-go/routes"
	"github.com/gjssss/soybean-admin-go/utils/cache"
	"github.com/gjssss/soybean-admin-go/utils/config"

	"github.com/gin-gonic/gin"

	// Swagger 相关导入
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// 导入生成的 swagger docs
	_ "github.com/gjssss/soybean-admin-go/docs"
)

// @title Soybean Admin Go API
// @version 1.0
// @description Soybean Admin Go 后台管理系统 API 文档

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description 请在此处输入Bearer Token，格式为: Bearer {token}
func main() {
	// 加载配置
	conf := config.Config{}
	err := conf.InitConfig()
	if err != nil {
		panic(err)
	}
	global.Config = &conf

	// 初始化数据库
	db := database.InitDB()
	global.DB = db
	rds := database.InitRedis()
	global.Redis = rds

	// 加载API缓存
	cache.ApiCache.Refresh()

	// 创建Gin实例
	router := gin.Default()

	// 初始化路由
	routes.Init(router)

	// 添加 Swagger API 文档路由
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1)))

	// 启动服务
	router.Run(":8080")
}
