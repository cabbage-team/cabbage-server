package model

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId   uuid.UUID `gorm:"type:binary(16);not null;comment:'uuid'" json:"userId"`
	Name     string    `gorm:"type:varchar(20);not null" json:"name"`
	Email    string    `gorm:"type:varchar(320);not null;unique" json:"email"`
	Password string    `gorm:"size:255;not null" json:"password"`
}

func (User) TableName() string {
	return fmt.Sprintf("%suser", viper.GetString("datasource.tableprefix"))
}
