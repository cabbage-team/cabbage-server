package model

import (
	"fmt"

	"github.com/spf13/viper"
)

type PostTag struct {
	PostId int64 `json:"postID" gorm:"column:post_id;index"`
	TagId  int64 `json:"tagId" gorm:"column:tag_id;index"`
}

func (PostTag) TableName() string {
	return fmt.Sprintf("%spost_tag", viper.GetString("datasource.tableprefix"))
}
