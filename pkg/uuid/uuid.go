package uuid

import (
	uuid "github.com/satori/go.uuid"
)

type UUIDGenerator interface {
	V4() string
}

type UUID struct {
}

func (u *UUID) V4() string {
	return uuid.NewV4().String()
}

func GenUUID(u UUIDGenerator) string {
	return u.V4()
}

func GetUUID() string {
	return GenUUID(&UUID{})
}
