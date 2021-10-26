package usecase

type IDiary interface {
	Create(req *CreateDiaryRequest) (*CreateDiaryResponse, error) //memo: rename to POST?
}

type CreateDiaryRequest struct {
}

type CreateDiaryResponse struct {
	Id string
}
