package system

import (
	"github.com/gjssss/soybean-admin-go/global"
	"github.com/gjssss/soybean-admin-go/models/system"
)

type ButtonRepository struct{}

func (c *ButtonRepository) GetButtons() ([]system.Button, error) {
	var buttons []system.Button
	err := global.DB.Find(&buttons).Error
	return buttons, err
}

func (c *ButtonRepository) GetButtonsByUserId(userId uint) ([]system.Button, error) {
	var buttons []system.Button
	err := global.DB.Raw(`
		SELECT b.* FROM buttons b
		JOIN role_buttons rb ON b.id = rb.button_id
		JOIN user_roles ur ON rb.role_id = ur.role_id
		WHERE ur.user_id = ?
	`, userId).Scan(&buttons).Error
	return buttons, err
}

func (c *ButtonRepository) GetButtonsByRoleId(roleId uint) ([]system.Button, error) {
	var buttons []system.Button
	err := global.DB.Raw(`
		SELECT b.* FROM buttons b
		JOIN role_buttons rb ON b.id = rb.button_id
		WHERE rb.role_id = ?
	`, roleId).Scan(&buttons).Error
	return buttons, err
}

func (c *ButtonRepository) CreateButton(button *system.Button) error {
	return global.DB.Create(button).Error
}

func (c *ButtonRepository) UpdateButton(button *system.Button) error {
	return global.DB.Model(&system.Button{}).Where("id = ?", button.ID).Updates(button).Error
}

func (c *ButtonRepository) DeleteButton(id uint) error {
	return global.DB.Delete(&system.Button{}, id).Error
}

func (c *ButtonRepository) BatchDeleteButton(ids []uint) error {
	return global.DB.Delete(&system.Button{}, ids).Error
}

func (c *ButtonRepository) IsCodeExist(code string, excludeID ...uint) (bool, error) {
	var count int64
	query := global.DB.Model(&system.Button{}).Where("code = ?", code)

	if len(excludeID) > 0 && excludeID[0] > 0 {
		query = query.Where("id != ?", excludeID[0])
	}

	err := query.Count(&count).Error
	return count > 0, err
}
