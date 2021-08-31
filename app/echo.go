package app

import (
	"strconv"

	"github.com/HideBa/notion-diary-auto/adapter/controller"
	"github.com/HideBa/notion-diary-auto/infrastructure/notion"
	"github.com/HideBa/notion-diary-auto/infrastructure/router"
	"github.com/HideBa/notion-diary-auto/interactor"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func NewEcho(config *Config) {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	gw := notion.NewNotionDiary()
	uc := interactor.NewDiary(&gw)
	con := controller.NewController(&uc)
	apiV1 := e.Group("/api/v1")
	router.Api(apiV1, con)

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
