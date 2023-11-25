package dao

import (
	"cabbage-server/db"
	"cabbage-server/model"
)

func GetUserByEmail(email string) (*model.User, error) {
	var DB = db.GetDB()
	var user model.User
	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
