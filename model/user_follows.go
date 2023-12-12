package model

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// UID -> Follow : Ship
type UserFollows struct {
	*gorm.Model
	// 用户id
	UID int64 `json:"uid" gorm:"index;column:uid"`
	// 目标用户id
	Follow int64 `json:"follower" gorm:"index:follow;column:follower"`
	// 关系 -1 拉黑 0 正常 1 关注
	Ship string `json:"ship" gorm:"column:ship;type:enum('F','N','D');default:0"`
}

func (UserFollows) TableName() string {
	return fmt.Sprintf("%suser_follows", viper.GetString("datasource.tableprefix"))
}
