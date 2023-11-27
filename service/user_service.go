package service

import (
	"cabbage-server/dao"
	"cabbage-server/internal/serverr"
	"cabbage-server/model"
	"errors"

	"gorm.io/gorm"
)

// CreateAccount 创建新用户服务
func CreateAccount(user *model.User) {
	dao.CreateAccount(user)
}

// GetUserProfile 获取用户信息服务
func GetUserProfile(email string) (*model.User, error) {
	user, err := dao.GetUserProfile(email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, serverr.RecordNotFoundError
		} else {
			return nil, serverr.InernalError
		}
	}
	return user, nil
}
