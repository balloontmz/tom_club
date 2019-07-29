package worker

import (
	"tom_club/app/model"
	"tom_club/app/server/assistant"
)

var (
	goodsChan = make(chan *assistant.GoodsResponse, 100) // 缓存中允许同时存在一百条记录
)

func runUpdateGoods() {
	go func() {
		//产生的迭代值为Channel中发送的值，它会一直迭代直到channel被关闭。
		for v := range goodsChan {
			go updateGoods(v)
		}
	}()
}

func updateGoods(g *assistant.GoodsResponse) {
	db := model.DB
	for _, val := range g.Data {
		goods := val.Goods
		db.Where(model.Goods{GoodsID: goods.GoodsID}).Assign(goods).FirstOrCreate(&goods) // 如果不存在，则创建。如果存在，则更新改记录的全部值
	}
}
