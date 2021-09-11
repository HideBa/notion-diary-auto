package interactor

import (
	"time"

	"github.com/HideBa/notion-diary-auto/gateway"
	"github.com/HideBa/notion-diary-auto/usecase"
)

type Diary struct {
	DiaryGateway   gateway.Diary
	WeatherGateway gateway.Weather
	NewsGateway    gateway.News
}

func NewDiary(d *gateway.Diary, w *gateway.Weather, n *gateway.News) usecase.IDiary {
	return &Diary{
		DiaryGateway:   *d,
		WeatherGateway: *w,
		NewsGateway:    *n,
	}
}

func (d *Diary) Create(req *usecase.CreateDiaryRequest) (*usecase.CreateDiaryResponse, error) {
	// sampleDiary := domain.NewDiary("hoge", "fuag")
	// d.WeatherGateway.TodaysWeather(time.Now(), domain.Location{Lat: 35.60904, Lng: 139.648021})
	d.NewsGateway.TodaysNews(time.Now())
	// d.DiaryGateway.AutoGenerate(sampleDiary)
	return &usecase.CreateDiaryResponse{
		Id: "hoge",
	}, nil
}
