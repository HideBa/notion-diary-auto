package interactor

import (
	"github.com/HideBa/notion-diary-auto/domain"
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
	sampleDiary := domain.NewDiary("hoge", "fuag")
	// w := d.WeatherGateway.TodaysWeather(time.Now(), domain.Location{Lat: 35.60904, Lng: 139.648021})
	// n := d.NewsGateway.TodaysNews(time.Now())
	// c := d.CalendarGateway.TodaysCalendar(time.Now())
	// diary := domain.NewDiary(w, n )

	d.DiaryGateway.Create(sampleDiary)
	return &usecase.CreateDiaryResponse{
		Id: "hoge",
	}, nil
}
