package system

type User struct {
	TimeRecord
	ID       uint   `gorm:"primaryKey" json:"id"`
	UserName string `gorm:"size:255;not null" json:"userName"`
	Password string `gorm:"size:255;not null" json:"password"`
	Roles    []Role `gorm:"many2many:user_roles" json:"roles,omitempty"`
}

type UserDTO struct {
	TimeRecord
	ID       uint     `json:"id"`
	UserName string   `json:"userName"`
	Roles    []string `json:"roles"`
	Buttons  []string `json:"buttons"`
}
