package service

import (
	"cabbage-server/dao"
	"cabbage-server/dto"
	"cabbage-server/internal"
	"errors"

	passlib "gopkg.in/hlandau/passlib.v1"
	"gorm.io/gorm"
)

func Register(account *dto.SignupDTO) error {
	hashString, err := passlib.Hash(account.Password)
	if err != nil {
		return internal.UserRegisterError
	}
	_,err = dao.CreateAccount(account.Name, hashString, account.Email)
	if err != nil {
		return internal.UserRegisterError
	}
	return nil
}

func Login(account *dto.LoginDTO) error {
	user, err := dao.FindUserByEmail(account.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return internal.UserLoginFail
		}
		return internal.UserNotExistsErr
	}
	newHash, err := passlib.Verify(account.Password, user.Password)
	if err != nil && newHash != "" {
		return internal.UserPasswordNotCorrectErr
	}
	return nil
}
