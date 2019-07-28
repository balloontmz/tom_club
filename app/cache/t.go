package cache

import (
	"tom_club/app/server/assistant"

	"github.com/labstack/gommon/log"
)

//GoodsCache 定义一个用于保存全局goods缓存的变量
// 此处需要注意必须采用 make 进行变量的初始化，不然该变量将是 nil 不可用状态，当然也可手动定义set函数
var GoodsCache = make(map[string]*assistant.GoodsResponse)

//GetTest 获取指定缓存
func GetTest() *assistant.GoodsResponse {
	goods := GoodsCache["test"]
	log.Print("获取到的缓存中goods为：", goods)
	return goods
}

//SetTest 设置指定键的全局缓存
func SetTest(key string, goods *assistant.GoodsResponse) error {
	GoodsCache[key] = goods
	log.Print("设置缓存变量")
	return nil
}
