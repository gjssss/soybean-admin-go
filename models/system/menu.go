package system

type Menu struct {
	ID        uint   `gorm:"primaryKey"  json:"id"`
	ParentID  *uint  `gorm:"index" json:"-"`
	Name      string `gorm:"unique;not null" json:"name"`
	Path      string `gorm:"unique;not null" json:"path"`
	Component string `gorm:"size:255" json:"component"`
	Meta      Meta   `json:"meta"`
	Children  []Menu `gorm:"foreignKey:ParentID" json:"children,omitempty"`
}

type Meta struct {
	ID               uint     `gorm:"primaryKey" json:"id"`
	MenuID           uint     `json:"-"`
	Title            string   `gorm:"not null" json:"title"`
	I18nKey          string   `json:"i18nKey,omitempty"`
	Order            int      `json:"order,omitempty"`
	KeepAlive        bool     `json:"keepAlive,omitempty"`
	Constant         bool     `json:"constant"`
	Icon             string   `json:"icon,omitempty"`
	LocalIcon        string   `json:"localIcon,omitempty"`
	Href             string   `json:"href,omitempty"`
	HideInMenu       bool     `json:"hideInMenu,omitempty"`
	ActiveMenu       string   `json:"activeMenu,omitempty"`
	MultiTab         bool     `json:"multiTab,omitempty"`
	FixedIndexInTabs bool     `json:"fixedIndexInTabs,omitempty"`
	Query            []Query  `json:"query"`
	Buttons          []Button `gorm:"many2many:meta_buttons" json:"buttons"`
}

type Query struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	MetaID uint   `json:"-"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}
