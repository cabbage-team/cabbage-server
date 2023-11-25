package service

import (
	"cabbage-server/dao"
	"cabbage-server/model"
)

func GetUserByEmail(email string) (*model.User, error) {
	return dao.GetUserByEmail(email)
}
