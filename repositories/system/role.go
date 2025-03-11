package system

import (
	"github.com/gjssss/soybean-admin-go/global"
	"github.com/gjssss/soybean-admin-go/models/common"
	"github.com/gjssss/soybean-admin-go/models/system"
)

type RoleRepository struct{}

func (c *RoleRepository) GetRoles(page common.PaginationParam) ([]system.Role, int64, error) {
	var roles []system.Role
	var count int64
	err := global.DB.Offset(page.PageSize * (page.Current - 1)).Limit(page.PageSize).Find(&roles).Error
	if err != nil {
		return nil, count, err
	}
	err = global.DB.Model(&system.Role{}).Count(&count).Error
	if err != nil {
		return nil, count, err
	}
	return roles, count, err
}

func (c *RoleRepository) GetAllRoles() ([]system.Role, error) {
	var roles []system.Role
	err := global.DB.Find(&roles).Error
	return roles, err
}
