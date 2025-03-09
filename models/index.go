package models

import (
	"github.com/gjssss/soybean-admin-go/models/system"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	system.AutoMigrate(db)
}

func InitDatabase(db *gorm.DB) {
	system.InitDatabase(db)
}
