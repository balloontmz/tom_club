package controller

import (
	"net/http"
	"strconv"

	"tom_club/app/cusresponse"
	"tom_club/app/model"
	"tom_club/app/server"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

//GetGoods 拉取商品
func GetGoods(c echo.Context) (err error) {
	page := c.QueryParam("page")
	size := c.QueryParam("size")
	pageInt, err := strconv.Atoi(page)
	sizeInt, err := strconv.Atoi(size)

	p := model.Params{Page: pageInt, Size: sizeInt}

	goods := model.GetGoods(p)
	return c.JSON(http.StatusOK, cusresponse.ResponseFmt{Ret: 1, Msg: "", Data: goods})
}

//Test 测试
func Test(c echo.Context) (err error) {
	goods, e := server.GetGoods(1)
	if err != nil {
		log.Print("拉取数据出错", e)
	}
	db := model.DB
	for _, val := range goods.Data {
		goods := val.Goods
		db.Where(model.Goods{GoodsID: goods.GoodsID}).Assign(goods).FirstOrCreate(&goods) // 如果不存在，则创建。如果存在，则更新改记录的全部值
	}
	return c.JSON(http.StatusOK, "拉取成功")
}
