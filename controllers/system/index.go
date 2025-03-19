package system

import (
	"github.com/gjssss/soybean-admin-go/services"
)

var (
	SystemService = services.System
)

type Group struct {
	User   UserController
	Menu   MenuController
	Role   RoleController
	Button ButtonController
	Api    ApiController
	Upload UploadController
}
