package domain

type UserRepository interface {
	GetUserByID(id string) (*User, error)
	CreateUser(newUser *User) (*User, error)
	UpdateUser(newUser *User) error
	Deleteuser(id string) error
}
