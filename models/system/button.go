package system

type Button struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Code string `json:"code"`
	Desc string `json:"desc"`
}
