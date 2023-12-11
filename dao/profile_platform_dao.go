package dao

import (
	"cabbage-server/db"
	"cabbage-server/model"
)

// 新增社交平台
func AddNewPlatform(name string) (*model.ProfilePlatform, error) {
	platform := &model.ProfilePlatform{
		Platform: name,
	}
	err := db.DB.Model(&model.ProfilePlatform{}).Create(platform).Error
	if err != nil {
		return nil, err
	} else {
		return platform, nil
	}
}

// 获取所有社交平台
func GetAllPlatform() ([]string, error) {
	platformList := []string{}
	err := db.DB.Model(&model.ProfilePlatform{}).
		Select("platform").
		Find(&platformList).
		Error
	if err != nil {
		return nil, err
	} else {
		return platformList, nil
	}
}

// 根据社交平台名称查找
func FindPlatformByName(name string) (*model.ProfilePlatform, error) {
	platform := &model.ProfilePlatform{}
	err := db.DB.Model(&model.ProfilePlatform{}).
		Omit("created_at", "deleted_at", "updated_at").
		Where("platform = ?", name).
		First(platform).
		Error
	if err != nil {
		return nil, err
	}
	return platform, nil
}
