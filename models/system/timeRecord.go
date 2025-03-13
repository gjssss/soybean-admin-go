package system

import (
	"time"

	"gorm.io/gorm"
)

type TimeRecord struct {
	CreatedAt time.Time      `json:"createTime"`
	UpdatedAt time.Time      `json:"updateTime"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
