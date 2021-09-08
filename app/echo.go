package app

import (
	"github.com/HideBa/notion-diary-auto/adapter/controller"
	"github.com/HideBa/notion-diary-auto/infrastructure/notion"
	"github.com/HideBa/notion-diary-auto/infrastructure/router"
	"github.com/HideBa/notion-diary-auto/infrastructure/weather"
	"github.com/HideBa/notion-diary-auto/interactor"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func NewEcho(config *Config) {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	ngw := notion.NewNotionDiary(&config.Notion)
	wgw := weather.NewYahooWeather(&config.Yahoo)
	uc := interactor.NewDiary(&ngw, &wgw)
	con := controller.NewController(uc)
	apiV1 := e.Group("/api/v1")
	router.Api(apiV1, con)

	if config.DebugMode == true {
		e.Logger.SetLevel(log.ERROR)
	}

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	e.Logger.Fatal(e.Start(":" + (config.App.Port)))
}
