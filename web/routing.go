package web

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"sync-mysql-data-server/pkg/config"
	"sync-mysql-data-server/web/middlewares"
	"sync-mysql-data-server/web/canal"
)

var conf = config.Config

func version(c echo.Context) error {
	return c.String(http.StatusOK, conf.GetString("version"))
}

func InitCanal() error {
	return canal.Init()
}

func SetupRoutes(router *echo.Echo) {
	// Middleware
	router.Use(middleware.Logger())
	setupRecover(router)
	router.GET("/", version)
	{
		router.Group("v1")
	}
}

func setupRecover(router *echo.Echo) {
	recoverMiddleware := middlewares.RecoverWithConfig(middlewares.RecoverConfig{
		StackSize: 10 << 10, // 10KB
	})
	router.Use(recoverMiddleware)
}
