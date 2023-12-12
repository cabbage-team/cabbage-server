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
	UID      int64 `json:"uid" gorm:"column:uid;index"`
	// 目标用户id
	Follow int64 `json:"follower" gorm:"column:follower;index"`
	// 关系 -1 拉黑 0 正常 1 关注
	Ship     int   `json:"ship" gorm:"column:ship;type:enum(-1,0,1);default:0"`
}

func (UserFollows) TableName() string {
	return fmt.Sprintf("%suser_follows", viper.GetString("datasource.tableprefix"))
}
