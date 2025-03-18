package system

import (
	"github.com/gjssss/soybean-admin-go/models/system"
)

type ApiService struct{}

func (s *ApiService) CreateApi(api *system.Api) error {
	return SystemRepositories.Api.CreateApi(api)
}

func (s *ApiService) UpdateApi(api *system.Api) error {
	return SystemRepositories.Api.UpdateApi(api)
}

func (s *ApiService) DeleteApi(id uint) error {
	return SystemRepositories.Api.DeleteApi(id)
}

func (s *ApiService) DeleteApiBatch(ids []uint) error {
	return SystemRepositories.Api.DeleteApiBatch(ids)
}

func (s *ApiService) GetApisByRoleID(roleID uint) ([]system.Api, error) {
	return SystemRepositories.Api.GetApisByRoleID(roleID)
}

func (s *ApiService) GetAllApis() ([]system.Api, error) {
	return SystemRepositories.Api.GetAllApis()
}

func (s *ApiService) GetRoleApis() ([]map[string]interface{}, error) {
	return SystemRepositories.Api.GetRoleApis()
}

func (s *ApiService) UpdateRoleApi(roleID uint, apiIDs []uint) error {
	return SystemRepositories.Api.UpdateRoleApi(roleID, apiIDs)
}
