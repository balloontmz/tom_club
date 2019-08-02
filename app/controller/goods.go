package controller

import (
	"net/http"

	"tom_club/app/model"
	"tom_club/app/server"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

//TestGoods 测试拉取商品
func TestGoods(c echo.Context) (err error) {
	log.Print("test")

	t, err := server.GetGoods(1)

	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusOK, "有bug")
	}

	// client := reqcli.NewClient()

	// t, _ := client.Get()

	log.Print(((t.Data[0]).Goods))
	db := model.DB

	for _, v := range t.Data {
		goods := v.Goods
		db.Where(model.Goods{GoodsID: goods.GoodsID}).Assign(goods).FirstOrCreate(&goods) // 如果不存在，则创建。如果存在，则更新改记录的全部值
		// log.Print("打印一下更新完的记录", goods)
		// db.Create(&goods)
	}

	return c.JSON(http.StatusOK, "test")
}
