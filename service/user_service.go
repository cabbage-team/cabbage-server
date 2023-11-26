package service

import (
	"cabbage-server/dao"
	"cabbage-server/model"
)

// CreateAccount 创建新用户服务
func CreateAccount(user *model.User) {
	dao.CreateAccount(user)
}

// GetUserProfile 获取用户信息服务
func GetUserProfile(email string) (*model.User, error) {
	return dao.GetUserProfile(email)
}
