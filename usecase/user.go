package usecase

type IUser interface {
	Create(req *CreateUserRequest) (*CreateUserResponse, error)
	Update(req *UpdateUserRequest) (*UpdateUserResponse, error)
	Delete(req *DeleteUserRequest) (*DeleteUserResponse, error)
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
