package main

import (
	"fmt"
	"github.com/labstack/echo/v4"

	"sync-mysql-data-server/pkg/config"
	"sync-mysql-data-server/web"
)

var conf = config.Config

func initServer() *echo.Echo {
	// Echo instance
	e := echo.New()
	// Routes
	web.SetupRoutes(e)
	go func() {
		err := web.InitCanal()
		if err != nil {
			panic(err)
		}
	}()
	return e
}

func main() {
	e := initServer()
	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", conf.GetString("port"))))
}