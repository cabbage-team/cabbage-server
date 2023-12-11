package model

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ProfilePlatform struct {
	*gorm.Model
	Platform string `json:"platform" gorm:"platform"`
}


func (ProfilePlatform) TableName() string {
	return fmt.Sprintf("%suser_profile_platform", viper.GetString("datasource.tableprefix"))
}
