package domain

import "github.com/HideBa/notion-diary-auto/pkg/uuid"

type User struct {
	id       string
	username string
}

func NewUser(username string) *User {
	id := uuid.GetUUID()
	return NewUserWithId(id, username)
}

func NewUserWithId(id string, username string) *User {
	return &User{
		id:       id,
		username: username,
	}
}

func UpdateUser(id string, username string) *User {
	return &User{
		id:       id,
		username: username,
	}
}

func (u *User) ID() string {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) SetUsername(un string) {
	u.username = un
}
