package dao

import (
	"cabbage-server/db"
	"cabbage-server/model"
)

// CreateAccount 创建用户数据库操作
func CreateAccount(user *model.User) error {
	err := db.DB.Model(&model.User{}).Create(user).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

// GetUserProfile 获取用户信息数据库操作
func FindUserByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := db.DB.Model(&model.User{}).Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func FindUserByName(name string) (*model.User, error) {
	user := &model.User{}
	err := db.DB.Model(&model.User{}).
		Where("name = ?", name).
		First(user).
		Error
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
