package system

import (
	"time"
)

type Role struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"createTime"`
	UpdatedAt time.Time `json:"updateTime"`
	DeletedAt time.Time `json:"deleteTime"`
	RoleName  string    `gorm:"size:255;not null" json:"roleName" binding:"required"`
	RoleDesc  string    `gorm:"size:255" json:"roleDesc"`
	Menus     []Menu    `gorm:"many2many:role_menus" json:"menu,omitempty"`
	Buttons   []Button  `gorm:"many2many:role_buttons" json:"buttons,omitempty"`
}
