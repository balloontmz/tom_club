package controller

import (
	"net/http"

	"tom_club/app/reqcli"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

//TestGoods 测试拉取商品
func TestGoods(c echo.Context) (err error) {
	log.Print("test")
	client := reqcli.NewClient()

	t, _ := client.Get()

	log.Print(t)

	return c.JSON(http.StatusOK, "test")
}
