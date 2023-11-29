package model

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	TagId       uuid.UUID `gorm:"type:binary(16);not null;comment:'uuid'"`
	TagName     string    `gorm:"type:varchar(20);not null;comment:'tag name'"`
	Description string    `gorm:"type:varchar(255);comment:'tag description'"`
	UsageCount  int       `gorm:"type:int;default:0;not null;comment:'usage count'"`
}

func (Tag) TableName() string {
	return fmt.Sprintf("%stag", viper.GetString("datasource.tableprefix"))
}
