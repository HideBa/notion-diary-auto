package controller

import (
	"github.com/HideBa/notion-diary-auto/interactor"
)

type Controller struct {
	DiaryController *DiaryController
	UserController  *UserController
}

type OuterController struct {
	OAuth *OAuthController
}

func NewInternalController(i *interactor.Interactor) *Controller {
	return &Controller{
		DiaryController: NewDiaryController(i.Diary),
		UserController:  NewUserController(i.User),
	}
}

//TODO: should not have 2 controllers
func NewOuterController(r *NotionRawConfig) *OuterController {
	return &OuterController{
		OAuth: NewOAuthController(r),
	}
}
