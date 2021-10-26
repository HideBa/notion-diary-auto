package router

import (
	"github.com/HideBa/notion-diary-auto/adapter/controller"
	"github.com/labstack/echo/v4"
)

func Api(r *echo.Group, c *controller.Controller, oc *controller.OuterController) {
	r.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, "pong")
	})
	r.POST("/test", func(c echo.Context) error {
		return c.JSON(200, "test")
	})
	r.GET("/diaries", c.DiaryController.Fetch)
	r.POST("/diaries", c.DiaryController.Create)
	r.POST("/notion/authorize", oc.OAuth.Notion.GetAuthCode)
	r.GET("/notion/callback", oc.OAuth.Notion.GetToken)
}
