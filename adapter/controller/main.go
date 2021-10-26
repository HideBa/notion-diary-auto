package controller

import "github.com/HideBa/notion-diary-auto/usecase"

type Controller struct {
	DiaryController *DiaryController
}

type OuterController struct {
	OAuth *OAuthController
}

func NewInternalController(i usecase.IDiary) *Controller {
	return &Controller{
		DiaryController: NewDiaryController(i),
	}
}

//TODO: should not have 2 controllers
func NewOuterController(r *NotionRawConfig) *OuterController {
	return &OuterController{
		OAuth: NewOAuthController(r),
	}
}
