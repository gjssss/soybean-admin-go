package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址（Docker 容器）
		Password: "",               // Redis 密码（如果有的话）
		DB:       0,                // 使用的数据库编号
	})
	rdb.FlushDB(context.Background())
	return rdb
}
