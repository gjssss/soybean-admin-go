package system

import (
	"errors"

	"github.com/gjssss/soybean-admin-go/models/system"
	"github.com/gjssss/soybean-admin-go/utils"
)

type UserService struct {
}

func (s *UserService) GetAllUsers() ([]system.User, error) {
	return SystemRepositories.User.FindAll()
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

func (s *UserService) UpdateUserPassword(user system.User) (system.User, error) {
	var err error
	user.Password, err = utils.EncodePassword(user.Password)
	if err != nil {
		return system.User{}, err
	}
	return SystemRepositories.User.Update(user)
}

func (s *UserService) DeleteUser(user system.User) error {
	return SystemRepositories.User.Delete(user)
}

func (s *UserService) GetUserById(id uint) (system.UserDTO, error) {
	user, err := SystemRepositories.User.FindById(id)
	if err != nil {
		return system.UserDTO{}, err
	}
	userDto := system.UserDTO{
		ID:       user.ID,
		UserName: user.UserName,
		Roles:    make([]string, len(user.Roles)),
		Buttons:  make([]string, len(user.Buttons)),
	}
	for i, role := range user.Roles {
		userDto.Roles[i] = role.RoleName
	}
	for i, button := range user.Buttons {
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
