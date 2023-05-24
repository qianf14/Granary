package main

import (
	"Granary/config"
	"Granary/dao"
	"Granary/routes"
	"github.com/spf13/viper"
)

// 程序入口
func main() {

	config.InitConfig()
	dao.InitDb()
	engine := routes.InitRoutes()
	engine.Run(viper.GetString("HttpPort"))

}
