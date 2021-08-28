package interactor

import (
	"github.com/HideBa/notion-diary-auto/gateway"
	"github.com/HideBa/notion-diary-auto/usecase"
)

type Diary struct {
	DiaryGateway gateway.Diary
}

func NewDiary(g gateway.Diary) usecase.IDiary {
	return &Diary{
		DiaryGateway: g,
	}
}

func (d *Diary) Create(req *usecase.CreateDiaryRequest) (*usecase.CreateDiaryResponse, error) {
	return &usecase.CreateDiaryResponse{
		Id: "hoge",
	}, nil
}
