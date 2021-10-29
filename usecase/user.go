package usecase

import "github.com/HideBa/notion-diary-auto/domain"

type IUser interface {
	FetchAll() (*GetUsersResponse, error)
	Create(req *CreateUserRequest) (*CreateUserResponse, error)
	Update(req *UpdateUserRequest) (*UpdateUserResponse, error)
	Delete(req *DeleteUserRequest) (*DeleteUserResponse, error)
}
type GetUsersResponse struct {
	users []domain.User //TODO: change here
}
type CreateUserRequest struct {
	Username string
}

type CreateUserResponse struct {
	Id       string
	Username string
}

type UpdateUserRequest struct {
	Id       string
	Username string
}

type UpdateUserResponse struct {
}

type DeleteUserRequest struct {
	Id string
}

type DeleteUserResponse struct {
}
