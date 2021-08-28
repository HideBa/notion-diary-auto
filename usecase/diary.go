package usecase

type IDiary interface {
	Create(req *CreateDiaryRequest) (*CreateDiaryResponse, error)
}

type CreateDiaryRequest struct {
}

type CreateDiaryResponse struct {
	Id string
}
