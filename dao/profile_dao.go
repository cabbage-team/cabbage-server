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

func FindProfileByUID(uid int64) ([]*model.UserProfile, error) {
	profile := []*model.UserProfile{}
	err := db.DB.Model(&model.UserProfile{}).
	Omit("created_at","deleted_at","updated_at").
		Where("user_id = ?", uid).
		Find(&profile).
		Error
	if err != nil {
		return nil, err
	} else {
		return profile, nil
	}
}

func FindSocials(uid int64, names []string) ([]*model.UserProfile, error) {
	socials := []*model.UserProfile{}
	err := db.DB.Model(&model.UserProfile{}).
		Select("platform", "uri").
		Where("id = ?", uid).
		Where("platform in (?)", names).
		Find(&socials).Error
	if err != nil {
		return nil, err
	} else {
		return socials, nil
	}
}

func FriendVisitProfile() ([]*model.UserProfile, error) {
	socials := []*model.UserProfile{}
	err := db.DB.Model(&model.UserProfile{}).
		Select("platform", "uri").
		Where("public = `Y`").
		Where("friend = `Y`").
		Find(&socials).
		Error
	if err != nil {
		return nil, err
	} else {
		return socials, nil
	}
}
