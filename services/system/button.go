package system

import (
	"github.com/gjssss/soybean-admin-go/models/system"
)

type ButtonService struct{}

func (s *ButtonService) GetButtonsByRoleId(roleId uint) ([]system.Button, error) {
	buttons, err := SystemRepositories.Button.GetButtonsByRoleId(roleId)
	return buttons, err
}

func (s *ButtonService) GetButtons() ([]system.Button, error) {
	buttons, err := SystemRepositories.Button.GetButtons()
	return buttons, err
}

func (s *ButtonService) GetButtonsByUserId(userId uint) ([]system.Button, error) {
	buttons, err := SystemRepositories.Button.GetButtonsByUserId(userId)
	return buttons, err
}
