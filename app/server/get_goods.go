package server

import (
	"context"
	"tom_club/app/server/assistant"
)

var (
	key = "b08fe2c272e19a9b"
)

//GetGoods test
func GetGoods() (*assistant.GoodsResponse, error) {
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
	resp, err := client.Goods(context.Background(), r)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
