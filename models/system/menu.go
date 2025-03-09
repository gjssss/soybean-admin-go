package system

type Menu struct {
	TimeRecord
	ID              uint        `gorm:"primaryKey;" json:"id"`
	Status          string      `json:"status"`
	ParentID        uint        `json:"parentId" gorm:"column:parent_id"` // 父菜单ID
	MenuType        string      `json:"menuType"`
	MenuName        string      `json:"menuName"`
	RouteName       string      `json:"routeName"`
	RoutePath       string      `json:"routePath"`
	Component       string      `json:"component"`
	Order           int         `json:"order"`
	I18nKey         string      `json:"i18nKey"`
	Icon            string      `json:"icon"`
	IconType        string      `json:"iconType"`
	MultiTab        bool        `json:"multiTab"`   // 是否支持多标签
	HideInMenu      bool        `json:"hideInMenu"` // 是否隐藏菜单
	ActiveMenu      string      `json:"activeMenu"` // 激活的菜单名
	Constant        bool        `json:"constant,omitempty"`
	FixedIndexInTab int         `json:"fixedIndexInTab,omitempty"`
	Href            string      `json:"href,omitempty"`
	KeepAlive       bool        `gorm:"default:false" json:"keepAlive,omitempty"`
	LocalIcon       string      `json:"localIcon,omitempty"`
	Children        []Menu      `gorm:"-" json:"children"`
	Query           []MenuQuery `json:"query,omitempty"`
	Button          []Button    `gorm:"many2many:menu_buttons" json:"buttons,omitempty"`
}

type MenuQuery struct {
	ID     uint   `gorm:"primaryKey;" json:"id"`
	MenuID uint   `json:"menuId"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}
