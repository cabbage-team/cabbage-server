package model

import (
	"fmt"

	"github.com/spf13/viper"
)

type PostTag struct {
	PostId uint `json:"postID" gorm:"column:post_id,index"`
	TagId  uint `json:"tagId" gorm:"column:post_id,index"`
}

func (PostTag) TableName() string {
	return fmt.Sprintf("%spost_tag", viper.GetString("datasource.tableprefix"))
}
