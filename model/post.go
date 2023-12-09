package model

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Post struct {
	*gorm.Model
	Title   string `json:"title" gorm:"column:title"`
	Content string `json:"content" gorm:"column:content"`
	Author  string `json:"author" gorm:"column:author"`
	Like int64 `json:"like" gorm:"column:like"`
	Diss int64 `json:"diss" gorm:"column:diss"`
	Share int64 `json:"share" gorm:"column:share"`
	Favorite int64 `json:"favorite" gorm:"column:favorite"`
	// Tags    []string `json:"tag" gorm:"column:tags"`
}

func (Post) TableName() string {
	return fmt.Sprintf("%spost", viper.GetString("datasource.tableprefix"))
}
