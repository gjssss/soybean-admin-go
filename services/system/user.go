package system

import (
	"errors"

	"github.com/gjssss/soybean-admin-go/models"
	"github.com/gjssss/soybean-admin-go/utils"
)

type UserService struct {
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return SystemRepositories.User.FindAll()
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	_, err := SystemRepositories.User.FindByUsername(user.UserName)
	if err == nil {
		return models.User{}, errors.New("用户名已存在")
	}
	user.Password, err = utils.EncodePassword(user.Password)
	if err != nil {
		return models.User{}, err
	}
	return SystemRepositories.User.Create(user)
}

func (s *UserService) UpdateUserPassword(user models.User) (models.User, error) {
	var err error
	user.Password, err = utils.EncodePassword(user.Password)
	if err != nil {
		return models.User{}, err
	}
	return SystemRepositories.User.Update(user)
}

func (s *UserService) DeleteUser(user models.User) error {
	return SystemRepositories.User.Delete(user)
}

func (s *UserService) GetUserById(id uint) (models.User, error) {
	return SystemRepositories.User.FindById(id)
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
