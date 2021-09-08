package interactor

import (
	"time"

	"github.com/HideBa/notion-diary-auto/domain"
	"github.com/HideBa/notion-diary-auto/gateway"
	"github.com/HideBa/notion-diary-auto/usecase"
)

type Diary struct {
	DiaryGateway   gateway.Diary
	WeatherGateway gateway.Weather
}

func NewDiary(d *gateway.Diary, w *gateway.Weather) usecase.IDiary {
	return &Diary{
		DiaryGateway:   *d,
		WeatherGateway: *w,
	}
}

func (d *Diary) Create(req *usecase.CreateDiaryRequest) (*usecase.CreateDiaryResponse, error) {
	// sampleDiary := domain.NewDiary("hoge", "fuag")
	d.WeatherGateway.TodaysWeather(time.Now(), domain.Location{Lat: 80, Lng: 130})
	// d.DiaryGateway.AutoGenerate(sampleDiary)
	return &usecase.CreateDiaryResponse{
		Id: "hoge",
	}, nil
}
