package main

import (
	conf "tom_club/app/config"
	routeregister "tom_club/app/routes"

	"github.com/labstack/gommon/log"
)

func main() {
	conf.InitConfig("config.ini") // 初始化全局配置文件

	conf.InitLog() // 此配置初步测试成功，如果新建文件夹需先创建文件夹

	log.Print("测试")

	router := routeregister.NewEngine() // 初始化路由

	// Listen and serve on 0.0.0.0:8080
	router.Logger.Fatal(router.Start(":8080"))
}
