package system

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createTime"`
	UpdatedAt time.Time      `json:"updateTime"`
	DeletedAt gorm.DeletedAt `json:"deleteTime"`
	RoleName  string         `gorm:"size:255;not null" json:"roleName"`
	RoleDesc  string         `gorm:"size:255" json:"roleDesc"`
	Menus     []Menu         `gorm:"many2many:role_menus" json:"menu,omitempty"`
}
