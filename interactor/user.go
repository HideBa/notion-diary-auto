package interactor

import (
	"github.com/HideBa/notion-diary-auto/domain"
	"github.com/HideBa/notion-diary-auto/usecase"
)

type User struct {
	UserRepository domain.UserRepository
}

func NewUser(repo *domain.UserRepository) usecase.IUser {
	return &User{
		UserRepository: *repo,
	}
}

func (u *User) FetchAll() (*usecase.GetUsersResponse, error) {
	return nil, nil
}

func (u *User) Create(req *usecase.CreateUserRequest) (res *usecase.CreateUserResponse, err error) {
	user, err := u.UserRepository.Create(domain.NewUser(req.Username))
	if err != nil {
		return nil, err
	}
	return &usecase.CreateUserResponse{Id: user.ID(), Username: user.Username()}, nil
}

func (u *User) Update(req *usecase.UpdateUserRequest) (res *usecase.UpdateUserResponse, err error) {
	_, err = u.UserRepository.GetByID(req.Id)
	if err != nil {
		return nil, err
	}
	err = u.UserRepository.Update(domain.UpdateUser(req.Id, req.Username))
	if err != nil {
		return nil, err
	}
	return &usecase.UpdateUserResponse{}, nil
}

func (u *User) Delete(req *usecase.DeleteUserRequest) (*usecase.DeleteUserResponse, error) {
	_, err := u.UserRepository.GetByID(req.Id)
	if err != nil {
		return nil, err
	}
	err = u.UserRepository.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
