package main

import (
	_ "cabbage-server/config"
	"cabbage-server/db"
	"cabbage-server/router"
	"fmt"
	"github.com/spf13/viper"
)

// @title           Cabbage server API
// @version         1.0
// @description     this is about novel swagger API
// @host            localhost:8080

func main() {
	// initialize service config
	db.InitDB()
	r := router.InitRouter()
	port := fmt.Sprintf(":%s", viper.GetString("server.port"))
	_ = r.Run(port)
}
