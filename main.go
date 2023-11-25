package main

import (
	"cabbage-server/config"
	"cabbage-server/db"
	"cabbage-server/router"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	// initialize service config
	config.InitConfig()
	db.InitDB()
	r := router.InitRouter()
	port := fmt.Sprintf(":%s", viper.GetString("server.port"))
	_ = r.Run(port)
}
