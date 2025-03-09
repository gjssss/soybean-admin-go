package system

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		User{},

		Menu{},
		Meta{},
		Button{},
		Query{},

		Role{},
	)
}

func InitDatabase(db *gorm.DB) {
	AutoMigrate(db)
	// m := db.Migrator()

	// User
	db.Create([]*User{
		{ID: 0, UserName: "admin", Password: EncodePassword("123123123")},
	})
}

func EncodePassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}
