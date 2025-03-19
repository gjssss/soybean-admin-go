package database

import (
	"context"

	"github.com/gjssss/soybean-admin-go/utils/config"
	"github.com/redis/go-redis/v9"
)

func InitRedis(conf config.RedisConfig) *redis.Client {
	var rdb *redis.Client
	if conf.Username != nil {
		rdb = redis.NewClient(&redis.Options{
			Addr:     conf.Host, // Redis 服务器地址（Docker 容器）
			DB:       conf.DB,   // 使用的数据库编号
			Username: *conf.Username,
			Password: *conf.Password,
		})
	} else {
		rdb = redis.NewClient(&redis.Options{
			Addr: conf.Host, // Redis 服务器地址（Docker 容器）
			DB:   conf.DB,   // 使用的数据库编号
		})
	}
	rdb.FlushDB(context.Background())
	return rdb
}
