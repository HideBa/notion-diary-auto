package interactor

import (
	"github.com/HideBa/notion-diary-auto/domain"
	"github.com/HideBa/notion-diary-auto/usecase"
)

type User struct {
	UserRepository domain.UserRepository
}

func NewUser(repo domain.UserRepository) *User {
	return &User{
		UserRepository: repo,
	}
}

func (u *User) Create(req *usecase.CreateUserRequest) (res *usecase.CreateUserResponse, err error) {
	user, err := u.UserRepository.CreateUser(domain.NewUser(req.Username))
	if err != nil {
		return nil, err
	}
	return &usecase.CreateUserResponse{Id: user.ID(), Username: user.Username()}, nil
}

func (u *User) Update(req *usecase.UpdateUserRequest) (res *usecase.UpdateUserResponse, err error) {
	_, err = u.UserRepository.GetUserByID(req.Id)
	if err != nil {
		return nil, err
	}
	err = u.UserRepository.UpdateUser(domain.UpdateUser(req.Id, req.Username))
	if err != nil {
		return nil, err
	}
	return &usecase.UpdateUserResponse{}, nil
}

func (u *User) Delete(req *usecase.DeleteUserRequest) error {
	_, err := u.UserRepository.GetUserByID(req.Id)
	if err != nil {
		return err
	}
	err = u.UserRepository.Deleteuser(req.Id)
	if err != nil {
		return err
	}
	return nil
}
