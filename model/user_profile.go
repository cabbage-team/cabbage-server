package model

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type UserProfile struct {
	*gorm.Model
	UserId   uint   `json:"userId" gorm:"column:user_id;index"`
	Platform string `json:"platform" gorm:"column:platform"`
	URI string `json:"uri" gorm:"column:uri"`
	OnlySelf string `json:"self" gorm:"column:only_self;type:enum('Y','N');default:'N'"`
	OnlyFriend string `json:"friend" gorm:"column:only_friend;type:enum('Y','N');default:'N'"`
	Public string `json:"public" gorm:"column:public;type:enum('Y','N');default:'Y'"`
}


func (UserProfile) TableName() string {
	return fmt.Sprintf("%suser_profile", viper.GetString("datasource.tableprefix"))
}
