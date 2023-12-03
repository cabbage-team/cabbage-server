package service

import (
	"cabbage-server/dao"
	"cabbage-server/model"
)

func ReadTags() *[]model.Tag {
	return dao.ReadTags()
}
