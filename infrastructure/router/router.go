package router

import (
	"github.com/HideBa/notion-diary-auto/adapter/controller"
	"github.com/labstack/echo/v4"
)

func Api(r *echo.Group, c *controller.Controller) {
	r.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, "pong")
	})
	r.GET("/diary", c.DiaryController.Fetch)
	r.POST("/diary", c.DiaryController.Create)
}
