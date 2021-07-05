package router

import (
	"github.com/labstack/echo/v4"
)

func Api(e *echo.Echo, r *echo.Group) {
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, "pong")
	})
}
