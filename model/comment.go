package model

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)


type Comment struct {
	*gorm.Model
	PostId int64 `json:"postId" gorm:"column:post_id"`
	CreateBy int64 `json:"createBy" gorm:"column:create_by"`
}


func (Comment) TableName() string {
	return fmt.Sprintf("%scomment", viper.GetString("datasource.tableprefix"))
}