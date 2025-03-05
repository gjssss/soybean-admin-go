package models

import (
	"time"

	"gorm.io/gorm"
)

type TimeRecord struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
