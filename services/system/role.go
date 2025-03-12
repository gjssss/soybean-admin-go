package system

import (
	"github.com/gjssss/soybean-admin-go/models/system"
	"github.com/gjssss/soybean-admin-go/utils"
)

type RoleService struct{}

func (s *RoleService) GetRoles(page utils.PaginationParam) ([]system.Role, int64, error) {
	return SystemRepositories.Role.GetRoles(page)
}

func (s *RoleService) GetAllRoles() ([]system.Role, error) {
	return SystemRepositories.Role.GetAllRoles()
}

func (s *RoleService) CreateRole(role *system.Role) error {
	return SystemRepositories.Role.CreateRole(role)
}

func (s *RoleService) UpdateRole(role *system.Role) error {
	return SystemRepositories.Role.UpdateRole(role)
}

func (s *RoleService) DeleteRole(id uint) error {
	return SystemRepositories.Role.DeleteRole(id)
}

func (s *RoleService) BatchDeleteRole(ids []uint) error {
	return SystemRepositories.Role.BatchDeleteRole(ids)
}

func (s *RoleService) UpdateRoleMenu(roleID uint, menuIDs []uint) error {
	return SystemRepositories.Role.UpdateRoleMenu(roleID, menuIDs)
}

func (s *RoleService) UpdateRoleButton(roleID uint, buttonIDs []uint) error {
	return SystemRepositories.Role.UpdateRoleButton(roleID, buttonIDs)
}
