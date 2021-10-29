package controller

import (
	"github.com/HideBa/notion-diary-auto/usecase"
	"github.com/labstack/echo/v4"
)

type RequestUserPost struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type RequestUserUpdate struct {
	Id       string `json:"id"`
	Username string `json:"id"`
}

type RequestUserDelete struct {
	Id string `json:"id"`
}

type UserResponse struct {
	Id       string `json:"id"`
	Username string `json:"id"`
}

type UserController struct {
	interactor usecase.IUser
}

func NewUserController(i usecase.IUser) *UserController {
	return &UserController{
		interactor: i,
	}
}

func (uc *UserController) Create(c echo.Context) error {
	req := usecase.CreateUserRequest{}
	uc.interactor.Create(&req)
	return c.JSON(200, "good")
}

func (uc *UserController) FetchAll(c echo.Context) error {
	uc.interactor.FetchAll()
	return c.JSON(200, "users")
}
