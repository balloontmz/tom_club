package controller

import (
	"net/http"
	"strconv"

	"tom_club/app/cusresponse"
	"tom_club/app/model"

	"github.com/labstack/echo"
)

//GetGoods 测试拉取商品
func GetGoods(c echo.Context) (err error) {
	page := c.QueryParam("page")
	size := c.QueryParam("size")
	pageInt, err := strconv.Atoi(page)
	sizeInt, err := strconv.Atoi(size)

	p := model.Params{Page: pageInt, Size: sizeInt}

	goods := model.GetGoods(p)
	return c.JSON(http.StatusOK, cusresponse.ResponseFmt{Ret: 1, Msg: "", Data: goods})
}

//GetGoods 测试拉取商品
func GetGoods(c echo.Context) (err error) {
	db := model.DB
	var goods []model.Goods
	db.Limit(3).Find(&goods)
	return c.JSON(http.StatusOK, cusresponse.ResponseFmt{Ret: 1, Msg: "test", Data: goods})
}
