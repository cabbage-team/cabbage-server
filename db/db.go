package db

import (
	"cabbage-server/model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"net/url"
)

var DB *gorm.DB

func InitDB(){
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: driverName,
		DSN:        args,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   viper.GetString("datasource.tableprefix"),
			SingularTable: true,
		},
	})
	if err != nil {
		panic("fail to connect database, err: " + err.Error())
	}
	initAutoMigrate(db)
	DB = db
}

func initAutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Comment{},
		&model.CommentOperator{},
		&model.Post{},
		&model.PostOperator{},
		&model.Tag{},
		&model.User{},
		&model.UserTag{},
		&model.UserProfile{},
		&model.ProfilePlatform{},
	)
	if err != nil {
	}
}