package model

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type PostOperator struct {
	*gorm.Model
	UserID int64 `json:"userId" gorm:"column:user_id;index"`
	PostID int64 `json:"postId" gorm:"column:post_id;index"`
	OPCode int   `json:"OPCode" gorm:"column:op_code"`
}

func (PostOperator) TableName() string {
	return fmt.Sprintf("%spost_operator", viper.GetString("datasource.tableprefix"))
}
