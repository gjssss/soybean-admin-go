package system

import (
	"errors"

	"github.com/gjssss/soybean-admin-go/models/system"
	"github.com/gjssss/soybean-admin-go/utils"
)

type UserService struct{}

func (s *UserService) GetAllUsers(page utils.PaginationParam) ([]system.User, int64, error) {
	return SystemRepositories.User.FindAll(page)
}

func (s *UserService) CreateUser(user system.User) (system.User, error) {
	_, err := SystemRepositories.User.FindByUsername(user.UserName)
	if err == nil {
		return system.User{}, errors.New("用户名已存在")
	}
	user.Password, err = utils.EncodePassword(user.Password)
	if err != nil {
		return system.User{}, err
	}
	return SystemRepositories.User.Create(user)
}

func (s *UserService) UpdateUserPassword(user system.User) error {
	var err error
	user.Password, err = utils.EncodePassword(user.Password)
	if err != nil {
		return err
	}
	return SystemRepositories.User.UpdatePassword(user)
}

func (s *UserService) DeleteUser(user system.User) error {
	return SystemRepositories.User.Delete(user)
}

func (s *UserService) BatchDeleteUser(ids []uint) error {
	for _, id := range ids {
		err := s.DeleteUser(system.User{ID: id})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *UserService) GetUserById(id uint) (system.UserDTO, error) {
	user, err := SystemRepositories.User.FindById(id)
	if err != nil {
		return system.UserDTO{}, err
	}
	btns, err := SystemRepositories.Button.GetButtonsByUserId(id)
	if err != nil {
		return system.UserDTO{}, err
	}
	userDto := system.UserDTO{
		ID:       user.ID,
		UserName: user.UserName,
		Roles:    make([]string, len(user.Roles)),
		Buttons:  make([]string, len(btns)),
	}
	for i, role := range user.Roles {
		userDto.Roles[i] = role.RoleName
	}
	for i, button := range btns {
		userDto.Buttons[i] = button.Code
	}
	return userDto, nil
}

type Token struct {
	AccessToken  string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

func (s *UserService) Login(username string, password string) (Token, error) {
	user, err := SystemRepositories.User.FindByUsername(username)
	if err != nil {
		return Token{}, errors.New("用户不存在")
	}
	if !utils.CheckPassword(password, user.Password) {
		return Token{}, errors.New("密码错误")
	}
	access, refresh, err := utils.GenerateTokens(user.ID)
	if err != nil {
		return Token{}, errors.New("生成Token失败")
	}
	return Token{AccessToken: access, RefreshToken: refresh}, nil
}

func (s *UserService) Refresh(accessToken string, refreshToken string) (Token, error) {
	accessToken, refreshToken, err := utils.RefreshToken(accessToken, refreshToken)
	return Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, err
}

func (s *UserService) UpdateUserRoles(userID uint, roleIDs []uint) error {
	// 检查用户是否存在
	_, err := SystemRepositories.User.FindById(userID)
	if err != nil {
		return errors.New("用户不存在")
	}

	// 检查每个角色是否存在
	for _, roleID := range roleIDs {
		exists, err := SystemRepositories.Role.RoleExists(roleID)
		if err != nil || !exists {
			return errors.New("角色不存在或ID无效")
		}
	}

	// 更新用户的角色
	return SystemRepositories.User.UpdateUserRoles(userID, roleIDs)
}

// 添加检查用户名是否存在的方法
func (s *UserService) CheckUserNameExists(userName string) bool {
	_, err := SystemRepositories.User.FindByUsername(userName)
	return err == nil // 如果没有错误，说明找到了用户，用户名存在
}

// 获取用户的角色列表
func (s *UserService) GetUserRoles(userID uint) ([]system.Role, error) {
	// 获取用户信息，包括关联的角色
	user, err := SystemRepositories.User.FindById(userID)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	return user.Roles, nil
}
