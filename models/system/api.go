package system

type Api struct {
	ID uint `gorm:"primaryKey" json:"id"`
	TimeRecord
	Name   string `json:"name"`
	Path   string `json:"path"`
	Method string `json:"method"`
	Group  string `json:"group"`
	Role   []Role `gorm:"many2many:api_role" json:"role"`
}
