package server

import (
	"context"
	"tom_club/app/cache"
	"tom_club/app/server/assistant"
)

var (
	key = "b08fe2c272e19a9b"
)

//GetGoods test
func GetGoods(p int) (*assistant.GoodsResponse, error) {
	var client *assistant.Client
	var err error

	if key != "" {
		client, err = assistant.NewClient(assistant.WithAPIKey(key)) // 根据 key 创建 api 客户端
	}
	if err != nil {
		return nil, err
	}

	r := &assistant.GoodsRequest{
		APIKey: key,
	}

	if resp := cache.GetTest(); resp != nil { // 此处测试条件判断语句内进行赋值。。。注意被赋值的变量的生命周期只在条件判断语句中
		return resp, nil
	}

	resp, err := client.Goods(context.Background(), r)
	cache.SetTest("test", resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
