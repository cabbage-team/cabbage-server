package service

import (
	"cabbage-server/dao"
	"cabbage-server/dto"
	"cabbage-server/internal"

	passlib "gopkg.in/hlandau/passlib.v1"
)

func Register(account *dto.SignupDTO) error {
	hashString, err := passlib.Hash(account.Password)
	if err != nil {
		return internal.UserRegisterError
	}
	err = dao.CreateAccount(account.Name,hashString,account.Email)
	if err != nil {
		return internal.UserRegisterError
	}
	return nil
}

func Login() {

}
