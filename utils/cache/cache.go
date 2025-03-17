package cache

import (
	"github.com/gjssss/soybean-admin-go/database"
	"github.com/gjssss/soybean-admin-go/global"
	"github.com/gjssss/soybean-admin-go/utils/config"
)

func InitCache() {
	// 初始化缓存
	if global.Config.CacheEngine == config.CacheRedis {
		// 初始化 Redis
		rdb := database.InitRedis()
		global.Redis = rdb
	}
}
