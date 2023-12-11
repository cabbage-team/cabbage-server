package model

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Comment struct {
	*gorm.Model
	UserID   int64  `json:"userId" gorm:"column:user_id"`
	PostID   int64  `json:"postId" gorm:"column:post_id;index"`
	Parent   int64  `json:"parent" gorm:"column:parent;index"`
	Content  string `json:"content" gorm:"column:content"`
	Reply    int64  `json:"reply" gorm:"column:reply"`
	CreateBy int64  `json:"createBy" gorm:"column:create_by"`
	Like     int64  `json:"like" gorm:"column:like"`
	Diss     int64  `json:"diss" gorm:"column:diss"`
	Share    int64  `json:"share" gorm:"column:share"`
	Favorite int64  `json:"favorite" gorm:"column:favorite"`
}

func (Comment) TableName() string {
	return fmt.Sprintf("%scomment", viper.GetString("datasource.tableprefix"))
}
