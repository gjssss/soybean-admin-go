package database

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gjssss/soybean-admin-go/global"
	"github.com/gjssss/soybean-admin-go/models"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	config := global.Config.Db
	var db *gorm.DB
	var err error

	// 根据数据库类型选择不同的连接方式
	switch config.Type {
	case "mysql":
		// MySQL连接
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s",
			config.User, config.Password, config.Host, config.Port, config.Name, config.Timezone)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
	case "postgres", "postgresql":
		// PostgreSQL连接
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			config.Host, config.User, config.Password, config.Name, strconv.Itoa(config.Port), config.SslMode, config.Timezone)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
	default:
		panic("Unsupported database type: " + config.Type)
	}

	if err != nil {
		panic("Failed to connect database: " + err.Error())
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	models.AutoMigrate(db)

	if !global.Config.IsInit {
		models.InitDatabase(db)
		global.Config.IsInit = true
		global.Config.SaveConfig()
	}
	return db
}
