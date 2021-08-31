package controller

import "github.com/HideBa/notion-diary-auto/usecase"

type Controller struct {
	DiaryController *DiaryController
}

func NewController(i *usecase.IDiary) *Controller {
	return &Controller{
		DiaryController: NewDiaryController(i),
	}
}
