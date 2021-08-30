package controller

import "github.com/HideBa/notion-diary-auto/usecase"

type Controller struct {
	diaryController *DiaryController
}

func NewController(i *usecase.IDiary) *Controller {
	return &Controller{
		diaryController: NewDiaryController(i),
	}
}
