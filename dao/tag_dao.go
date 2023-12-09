package dao

import (
	"cabbage-server/db"
	"cabbage-server/model"
)

func ReadTags() (*[]model.Tag, error) {
	var tags []model.Tag
	err := db.DB.Model(&model.Tag{}).Find(&tags).Error
	if err != nil {
		return nil, err
	} else {
		return &tags, nil
	}
}
