package dao

import (
	"cabbage-server/db"
	"cabbage-server/model"
)

func FindProfileByUID(uid int64) (*model.UserProfile,error){
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
