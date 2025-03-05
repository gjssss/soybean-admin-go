package main

import (
	"github.com/gjssss/soybean-admin-go/config"
	"github.com/gjssss/soybean-admin-go/database"
	"github.com/gjssss/soybean-admin-go/global"
	"github.com/gjssss/soybean-admin-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	config := &config.Config{}
	config.Init()
	global.Config = config

	// 初始化数据库
	db := database.InitDB(config.DB)
	global.DB = db

	// 创建Gin实例
	router := gin.Default()

	// 初始化路由
	routes.Init(router)

	// 启动服务
	router.Run(":8080")
}
