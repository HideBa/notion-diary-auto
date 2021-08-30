package controller

import (
	"github.com/HideBa/notion-diary-auto/usecase"
	"github.com/labstack/echo/v4"
)

type RequestDiaryCreate struct {
}

type ResponseDiaryCreate struct {
	Id string
}

type DiaryController struct {
	interactor *usecase.IDiary
}

func NewDiaryController(i *usecase.IDiary) *DiaryController {
	return &DiaryController{
		interactor: i,
	}
}

func (d *DiaryController) Create(c echo.Context) error {
	return c.JSON(200, "helloo")
}
