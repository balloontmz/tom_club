package main

import (
	conf "tom_club/app/config"
	"tom_club/app/model"
	routeregister "tom_club/app/routes"
	"tom_club/app/worker"
)

func main() {
	conf.InitConfig("config.ini") // 初始化全局配置文件

	conf.InitLog() // 此配置初步测试成功，如果新建文件夹需先创建文件夹

	if _, err := model.InitDB(conf.DBConfig("config.ini")); err != nil { // 初始化数据库链接
		panic(err)
	}

	model.Migrate() // 数据库迁移

	worker.Run()

	router := routeregister.NewEngine() // 初始化路由

	// Listen and serve on 0.0.0.0:8080
	router.Logger.Fatal(router.Start(":9000"))
}
