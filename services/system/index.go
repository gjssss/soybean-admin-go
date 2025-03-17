package system

import "github.com/gjssss/soybean-admin-go/repositories"

var (
	SystemRepositories = repositories.System
)

type Group struct {
	User   UserService
	Menu   MenuService
	Role   RoleService
	Button ButtonService
	Api    ApiService
}
