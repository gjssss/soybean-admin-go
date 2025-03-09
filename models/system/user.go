package system

type User struct {
	TimeRecord
	ID       uint     `gorm:"primaryKey" json:"userId"`
	UserName string   `gorm:"size:255;not null" json:"userName"`
	Password string   `gorm:"size:255;not null" json:"password"`
	Roles    []Role   `gorm:"many2many:user_roles" json:"roles"`
	Buttons  []Button `gorm:"many2many:user_buttons" json:"buttons"`
}
