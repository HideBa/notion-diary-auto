package app

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func NewEcho() {
	e := echo.New()
	apiV1 := e.Group("/api/v1")

	config := NewConfig()

	if config.debugMode == true {
		e.Logger.SetLevel(log.ERROR)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(config.app.Port)))
}
