package main

import (
	"net/http"
	"time"

	"./routes"
	"./config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	e.Use(middleware.Logger())

	e.Any("/*", func(c echo.Context) (err error) {
		routes.APIRoutes().ServeHTTP(c.Response(), c.Request())
		return
	})
	serverConfig := &http.Server{
		Addr:         ":" + config.Port,
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	e.Logger.Fatal(e.StartServer(serverConfig))
}
