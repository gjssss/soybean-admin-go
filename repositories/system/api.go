package system

import (
	"github.com/gjssss/soybean-admin-go/global"
	"github.com/gjssss/soybean-admin-go/models/system"
)

type ApiRepository struct{}

func (r *ApiRepository) CreateApi(api *system.Api) error {
	return global.DB.Create(api).Error
}

func (r *ApiRepository) UpdateApi(api *system.Api) error {
	return global.DB.Save(api).Error
}

func (r *ApiRepository) DeleteApi(id uint) error {
	return global.DB.Delete(&system.Api{}, id).Error
}

func (r *ApiRepository) GetApisByRoleID(roleID uint) ([]system.Api, error) {
	var role system.Role
	var apis []system.Api
	err := global.DB.Preload("Apis").First(&role, roleID).Error
	if err != nil {
		return nil, err
	}
	apis = role.Apis
	return apis, nil
}

func (r *ApiRepository) GetAllApis() ([]system.Api, error) {
	var apis []system.Api
	err := global.DB.Find(&apis).Error
	return apis, err
}

func (r *ApiRepository) GetAllApisRoles() ([]system.Api, error) {
	var apis []system.Api
	err := global.DB.Preload("Roles").Find(&apis).Error
	return apis, err
}

func (r *ApiRepository) GetRoleApis() ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := global.DB.Table("role_api").Find(&results).Error
	return results, err
}

func (r *ApiRepository) UpdateRoleApi(roleID uint, apiIDs []uint) error {
	// 先清除角色的所有API权限
	role := system.Role{ID: roleID}
	if err := global.DB.Model(&role).Association("Apis").Clear(); err != nil {
		return err
	}

	// 如果没有新API要添加，则直接返回
	if len(apiIDs) == 0 {
		return nil
	}

	// 添加新API权限
	var apis []system.Api
	if err := global.DB.Find(&apis, apiIDs).Error; err != nil {
		return err
	}
	return global.DB.Model(&role).Association("Apis").Append(&apis)
}
