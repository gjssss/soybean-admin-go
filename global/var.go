package global

import (
	"github.com/gjssss/soybean-admin-go/config"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	// DB is a global variable that stores the database connection
	DB *gorm.DB
)
