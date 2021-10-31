package domain

import id "github.com/HideBa/notion-diary-auto/pkg/id"

type User struct {
	id       id.ID
	username string
}

func NewUser(username string) *User {
	id := id.NewID()
	return NewUserWithId(id, username)
}

func NewUserWithId(id id.ID, username string) *User {
	return &User{
		id:       id,
		username: username,
	}
}

func UpdateUser(id id.ID, username string) *User {
	return &User{
		id:       id,
		username: username,
	}
}

func (u *User) ID() id.ID {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) SetUsername(un string) {
	u.username = un
}
