package routeregister

import (
	"tom_club/app/controller"

	"github.com/labstack/echo"
)

//goodsRoutesRegister is a simple
func goodsRoutesRegister(router *echo.Echo) *echo.Echo {
	router.GET("/test", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	router.GET("/test-goods", controller.TestGoods)
	router.GET("/goods", controller.GetGoods)

	// router.GET("/json-test", controller.TestJSON)

	// email := router.Group("email")
	// email.Use(custommid.Auth)
	// {
	// 	email.POST("/send", controller.SendEmail)
	// }

	return router
}
