package model

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type UserTag struct {
	gorm.Model
	UserId       uuid.UUID `gorm:"type:binary(16);not null;comment:'uuid'" json:"userId"`
	TagId        uuid.UUID `gorm:"type:binary(16);not null;comment:'uuid'" json:"tagId"`
	Relationship byte      `gorm:"type:varchar(1);not null;comment:'1: 正常状态, 2: follow, 3: hide'" json:"relationship"`
}

func (UserTag) TableName() string {
	return fmt.Sprintf("%suser_tag", viper.GetString("datasource.tableprefix"))
}
