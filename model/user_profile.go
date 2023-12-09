package model

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type UserProfile struct {
	*gorm.Model
	UserId   uint   `json:"userId" gorm:"column:user_id;index"`
	Twitter  string `json:"twitter" gorm:"column:twitter"`
	Mastodon string `json:"mastodon" gorm:"column:mastodon"`
	Facebook string `json:"facebook" gorm:"column:facebook"`
	Youtobe  string `json:"youtobe" gorm:"column:youtobe"`
	Gmail    string `json:"gmail" gorm:"column:gmail"`
	Github   string `json:"github" gorm:"column:github"`
	Insgram  string `json:"insgram" gorm:"column:insgram"`
	Telegram string `json:"telegram" gorm:"column:telegram"`
}


func (UserProfile) TableName() string {
	return fmt.Sprintf("%suser_profile", viper.GetString("datasource.tableprefix"))
}
