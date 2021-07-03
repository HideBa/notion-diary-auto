package domain

type User struct {
	id       string
	username string
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
