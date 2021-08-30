package router

import (
	"github.com/labstack/echo/v4"
)

func Api(r *echo.Group) {
	r.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, "pong")
	})
	r.POST("/diary", 
}
