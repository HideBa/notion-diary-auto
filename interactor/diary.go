package interactor

import (
	"time"

	"github.com/HideBa/notion-diary-auto/gateway"
	"github.com/HideBa/notion-diary-auto/usecase"
)

type Diary struct {
	DiaryGateway    gateway.Diary
	WeatherGateway  gateway.Weather
	NewsGateway     gateway.News
	CalendarGateway gateway.Calendar
}

func NewDiary(d *gateway.Diary, w *gateway.Weather, n *gateway.News, c *gateway.Calendar) usecase.IDiary {
	return &Diary{
		DiaryGateway:    *d,
		WeatherGateway:  *w,
		NewsGateway:     *n,
		CalendarGateway: *c,
	}
}

func (d *Diary) Create(req *usecase.CreateDiaryRequest) (*usecase.CreateDiaryResponse, error) {
	// sampleDiary := domain.NewDiary("hoge", "fuag")
	// d.WeatherGateway.TodaysWeather(time.Now(), domain.Location{Lat: 35.60904, Lng: 139.648021})
	// d.NewsGateway.TodaysNews(time.Now())
	d.CalendarGateway.TodaysCalendar(time.Now())
	// d.DiaryGateway.AutoGenerate(sampleDiary)
	return &usecase.CreateDiaryResponse{
		Id: "hoge",
	}, nil
}
