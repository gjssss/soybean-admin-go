package global

import (
	"github.com/gjssss/soybean-admin-go/utils/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	// DB is a global variable that stores the database connection
	DB *gorm.DB
	// Redis is a global variable that stores the Redis connection
	Redis *redis.Client
)
