package system

import (
	"errors"

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

func (s *ButtonService) CreateButton(button *system.Button) error {
	// 检查按钮Code是否存在
	exists, err := SystemRepositories.Button.IsCodeExist(button.Code)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("按钮代码已存在")
	}

	return SystemRepositories.Button.CreateButton(button)
}

func (s *ButtonService) UpdateButton(button *system.Button) error {
	// 检查按钮Code是否与其他按钮冲突
	exists, err := SystemRepositories.Button.IsCodeExist(button.Code, button.ID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("按钮代码已存在")
	}

	return SystemRepositories.Button.UpdateButton(button)
}

func (s *ButtonService) DeleteButton(id uint) error {
	return SystemRepositories.Button.DeleteButton(id)
}

func (s *ButtonService) BatchDeleteButton(ids []uint) error {
	return SystemRepositories.Button.BatchDeleteButton(ids)
}
