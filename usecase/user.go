package usecase

import (
	"github.com/HideBa/notion-diary-auto/pkg/id"
)

type IUser interface {
	FetchAll() (*GetUsersResponse, error)
	Create(req *CreateUserRequest) (*CreateUserResponse, error)
	Update(req *UpdateUserRequest) (*UpdateUserResponse, error)
	Delete(req *DeleteUserRequest) (*DeleteUserResponse, error)
}
type GetUsersResponse struct {
	Users []UserResponse //TODO: change here
}

type UserResponse struct {
	Username string `json:"username"`
}
type CreateUserRequest struct {
	Username string
}

type CreateUserResponse struct {
	Id       id.ID
	Username string
}

type UpdateUserRequest struct {
	Id       id.ID
	Username string
}

type UpdateUserResponse struct {
}

type DeleteUserRequest struct {
	Id id.ID
}

type DeleteUserResponse struct {
}
