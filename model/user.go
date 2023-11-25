package model

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Email    string `gorm:"type:varchar(320);not null;unique"`
	Password string `gorm:"size:255;not null"`
}

func (User) TableName() string {
	return fmt.Sprintf("%suser", viper.GetString("datasource.tableprefix"))
}
