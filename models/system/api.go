package system

type Api struct {
	ID uint `gorm:"primaryKey" json:"id"`
	TimeRecord
	Path   string `json:"path"`
	Method string `json:"method"`
	Group  string `json:"group"`
	Roles  []Role `gorm:"many2many:role_api" json:"roles,omitempty"`
}
