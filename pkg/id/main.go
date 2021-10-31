package id

import "github.com/google/uuid"

type ID struct {
	id uuid.UUID
}

func (id *ID) GetID() uuid.UUID {
	return id.id
}

func NewID() ID {
	return ID{GenerateID()}
}

func GenerateID() uuid.UUID {
	return uuid.New()
}

func MustBeID(sID string) ID {
	parsedID, err := uuid.Parse(sID)
	if err != nil {
		panic("invalid id")
	}
	return ID{parsedID}
}

func FromID(sID string) (ID, error) {
	parsedID, err := uuid.Parse(sID)
	if err != nil {
		return ID{}, err
	}
	return ID{parsedID}, nil
}

func FromIDUUID(sID uuid.UUID) ID {
	return ID{sID}
}

func (id *ID) String() string {
	return id.GetID().String()
}
