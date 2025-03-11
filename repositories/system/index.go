package system

type Group struct {
	User UserRepository
	Menu MenuRepository
	Role RoleRepository
}
