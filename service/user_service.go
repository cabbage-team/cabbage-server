package service

import (
	"cabbage-server/dao"
	"cabbage-server/dto"
	"cabbage-server/internal"
	"cabbage-server/model"
	"errors"

	"gorm.io/gorm"
)

// CreateAccount 创建新用户服务
func CreateAccount(user *dto.SignupDTO) error {
	err := dao.CreateAccount(&model.User{
		Email: user.Email,
		Name: user.Name,
		Password: user.Password,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return internal.UserNotFoundError
		} else {
			return internal.InernalError
		}
	}
	return nil
}

// GetUserProfile 获取用户信息服务
func GetUserProfile(email string) (*model.User, error) {
	user, err := dao.GetUserProfile(email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, internal.RecordNotFoundError
		} else {
			return nil, internal.InernalError
		}
	}
	return user, nil
}
