package system

import (
	"github.com/gjssss/soybean-admin-go/services"
)

var (
	SystemService = services.System
)

type Group struct {
	User UserController
	Menu MenuController
}
