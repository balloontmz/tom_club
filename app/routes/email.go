package routeregister

import (
	"github.com/labstack/echo"
)

//EmailRoutesRegister is a simple
func emailRoutesRegister(router *echo.Echo) *echo.Echo {
	// router.GET("/test", func(c echo.Context) error {
	// 	return c.String(200, "pong")
	// })

	// router.GET("/sqs-test", controller.TestSQS)

	// router.GET("/json-test", controller.TestJSON)

	// email := router.Group("email")
	// email.Use(custommid.Auth)
	// {
	// 	email.POST("/send", controller.SendEmail)
	// }

	return router
}
