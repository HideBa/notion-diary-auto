package controller

import (
	"github.com/HideBa/notion-diary-auto/infrastructure/notion"
	"github.com/HideBa/notion-diary-auto/usecase"
	"github.com/labstack/echo/v4"
)

type RequestDiaryCreate struct {
}

type ResponseDiaryCreate struct {
	Id string
}

type DiaryController struct {
	interactor usecase.IDiary
}

func NewDiaryController(i usecase.IDiary) *DiaryController {
	return &DiaryController{
		interactor: i,
	}
}

func (dc *DiaryController) Create(c echo.Context) error {
	var r usecase.CreateDiaryRequest
	dc.interactor.Create(&r)
	return c.JSON(200, "helloo")
}

func (dc *DiaryController) Fetch(c echo.Context) error {
	return c.JSON(200, "fetch")
}

func (dc *DiaryController) ConnectNotion(c echo.Context) error {
	var r struct{}
	dc.interactor.ConnectNotion(&r)
	return c.JSON(200, "notion")
}

func (dc *DiaryController) CallbackNotion(c echo.Context) error {
	// var r struct{}
	notion.Callback()
	return nil
}
