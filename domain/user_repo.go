package domain

import "github.com/HideBa/notion-diary-auto/pkg/id"

type UserRepository interface {
	GetByID(id id.ID) (*User, error)
	GetAll() ([]User, error)
	Create(newUser *User) (*User, error)
	Update(newUser *User) error
	Delete(id id.ID) error
}
