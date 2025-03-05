package main

import (
	"github.com/gjssss/soybean-admin-go/config"
	"github.com/gjssss/soybean-admin-go/controllers"
	"github.com/gjssss/soybean-admin-go/database"
	"github.com/gjssss/soybean-admin-go/repositories"
	"github.com/gjssss/soybean-admin-go/routes"
	"github.com/gjssss/soybean-admin-go/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载环境变量
	config.LoadEnv()
	dbConfig := config.NewDBConfig()

	// 初始化数据库
	db := database.InitDB(dbConfig)

	// 初始化各层依赖
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// 创建Gin实例
	router := gin.Default()

	// 注册路由
	routes.SetupRoutes(router, userController)

	// 启动服务
	router.Run(":8080")
}
