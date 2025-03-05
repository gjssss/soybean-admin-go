package models

type User struct {
	TimeRecord
	ID       uint   `gorm:"primaryKey" json:"userId"`
	UserName string `gorm:"size:255;not null" json:"userName"`
}
