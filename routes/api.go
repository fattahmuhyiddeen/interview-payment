package routes

import (
	controller "../controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// APIRoutes are where routes are defined
func APIRoutes() *echo.Echo {
	api := echo.New()
	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	api.Use(middleware.Logger())
	api.Use(middleware.Recover())

	api.GET("/", controller.HomePage)
	api.GET("/timestamp", controller.Timestamp)
	api.GET("/payment_status", controller.PaymentStatus)

	return api
}
