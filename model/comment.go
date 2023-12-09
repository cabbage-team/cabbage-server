package model

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)


type Comment struct {
	*gorm.Model
	UserID int64 `json:"userId" gorm:"column:user_id"`
	PostId int64 `json:"postId" gorm:"column:post_id,index"`
	Parent int64 `json:"parent" gorm:"column:parent,index"`
	Content string `json:"content" gorm:"column:content"`
	CreateBy int64 `json:"createBy" gorm:"column:create_by"`
}


func (Comment) TableName() string {
	return fmt.Sprintf("%scomment", viper.GetString("datasource.tableprefix"))
}