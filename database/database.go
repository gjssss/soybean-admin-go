package database

import (
	"fmt"
	"time"

	"github.com/gjssss/soybean-admin-go/config"
	"github.com/gjssss/soybean-admin-go/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config *config.DBConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode, config.TimeZone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	models.AutoMigrate(db)

	return db
}
