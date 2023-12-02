package dao

import (
	"cabbage-server/db"
	"cabbage-server/model"
)

func ReadTags() *[]model.Tag {
	var DB = db.GetDB()
	var tags []model.Tag
	var sql = `SELECT * FROM cabbage_tag`
	DB.Raw(sql).Scan(&tags)
	return &tags
}
