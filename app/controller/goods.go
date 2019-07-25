package controller

import (
	"net/http"

	"tom_club/app/server"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

//TestGoods 测试拉取商品
func TestGoods(c echo.Context) (err error) {
	log.Print("test")

	t, err := server.GetGoods()

	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusOK, "有bug")
	}

	// client := reqcli.NewClient()

	// t, _ := client.Get()

	log.Print(t)

	return c.JSON(http.StatusOK, "test")
}
