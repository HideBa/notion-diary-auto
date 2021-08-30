package app

import (
	"strconv"

	"github.com/HideBa/notion-diary-auto/adapter/controller"
	"github.com/HideBa/notion-diary-auto/infrastructure/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/swaggo/swag/example/celler/controller"
)

func NewEcho(config *Config) {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	con = controller.NewController()
	apiV1 := e.Group("/api/v1")
	router.Api(apiV1)

	if config.debugMode == true {
		e.Logger.SetLevel(log.ERROR)
	}

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(config.app.Port)))
}
