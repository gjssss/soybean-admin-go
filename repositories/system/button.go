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
