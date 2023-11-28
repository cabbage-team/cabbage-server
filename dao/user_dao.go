package dao

import (
	"cabbage-server/db"
	"cabbage-server/model"
)

// CreateAccount 创建用户数据库操作
func CreateAccount(user *model.User) error {
	var DB = db.GetDB()
	if err := DB.Where("email = ?", user.Email).First(&user).Error; err != nil {
		return err
	}
	DB.Create(&user)
	return nil
}

// GetUserProfile 获取用户信息数据库操作
func GetUserProfile(email string) (*model.User, error) {
	var DB = db.GetDB()
	var user model.User
	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
