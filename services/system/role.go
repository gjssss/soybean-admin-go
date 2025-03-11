package system

import (
	"github.com/gjssss/soybean-admin-go/models/common"
	"github.com/gjssss/soybean-admin-go/models/system"
)

type RoleService struct{}

func (s *RoleService) GetRoles(page common.PaginationParam) ([]system.Role, int64, error) {
	return SystemRepositories.Role.GetRoles(page)
}
