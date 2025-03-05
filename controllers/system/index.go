package system

import (
	"github.com/gjssss/soybean-admin-go/services"
)

var (
	UserService = services.System.User
)

type Group struct {
	User UserController
}
