package usecase

type IDiary interface {
	Create(req *CreateDiaryRequest) (*CreateDiaryResponse, error) //memo: rename to POST?
	// ConnectNotion(req *struct{}) (*struct{}, error)
}

type CreateDiaryRequest struct {
}

type CreateDiaryResponse struct {
	Id     string
	連携先URL string
}
