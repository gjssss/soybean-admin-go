package database

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gjssss/soybean-admin-go/global"
	"github.com/gjssss/soybean-admin-go/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	config := global.Config.Db
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Host, config.User, config.Password, config.Name, strconv.Itoa(config.Port), config.SslMode, config.Timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("Failed to connect database")
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
