package domain

type UserRepository interface {
	GetByID(id string) (*User, error)
	Create(newUser *User) (*User, error)
	Update(newUser *User) error
	Delete(id string) error
}
