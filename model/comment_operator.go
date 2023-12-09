package model

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type CommentOperator struct {
	*gorm.Model
	UserID    int64 `json:"userid" gorm:"column:user_id;index"`
	CommentID int64 `json:"commentId" gorm:"column:comment_id;index"`
	OPCode    int   `json:"OPCode" gorm:"column:op_code"`
}

func (CommentOperator) TableName() string {
	return fmt.Sprintf("%scomment_operator", viper.GetString("datasource.tableprefix"))
}