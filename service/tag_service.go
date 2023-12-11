package service

import (
	"cabbage-server/dao"
	"cabbage-server/internal"
	"cabbage-server/model"
	"errors"

	"gorm.io/gorm"
)

func ReadTags(page, size int) (*[]model.Tag, error) {
	tags, err := dao.ReadTags(page,size)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, internal.InernalError
		}
		return nil, internal.RecordNotFoundError
	}
	return tags, nil
}
