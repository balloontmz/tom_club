package conf

import (
	"tom_club/app/model"
)

//DBConfig 返回文件中的数据库
func DBConfig(path string) *model.DBConfig {
	configMap, _ := GetConfig(path)
	config := &model.DBConfig{
		Username: configMap["DB_USERNAME"],
		Password: configMap["DB_PASSWORD"],
		Hostname: configMap["DB_HOSTNAME"],
		Port:     configMap["DB_PORT"],
		Database: configMap["DB_DATABASE"],
	}
	return config
}
