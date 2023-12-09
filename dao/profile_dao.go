package dao

import (
	"cabbage-server/db"
	"cabbage-server/model"
)

func CreateProfile(profile *model.UserProfile) (*model.UserProfile, error) {
	err := db.DB.Model(&model.UserProfile{}).Create(profile).Error
	if err != nil {
		return nil, err
	} else {
		return profile, err
	}
}

func FindProfileByUID(uid int64) (*model.UserProfile, error) {
	profile := &model.UserProfile{}
	err := db.DB.Model(&model.UserProfile{}).
		Where("user_id = ?", uid).
		First(profile).
		Error
	if err != nil {
		return nil, err
	} else {
		return profile, nil
	}
}
