package interactor

import (
	"github.com/HideBa/notion-diary-auto/gateway"
	"github.com/HideBa/notion-diary-auto/infrastructure/notion"
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
	n := PersonalDiary
	// w := d.WeatherGateway.TodaysWeather(time.Now(), domain.Location{Lat: 35.60904, Lng: 139.648021})
	// n := d.NewsGateway.TodaysNews(time.Now())
	// c := d.CalendarGateway.TodaysCalendar(time.Now())
	// diary := domain.NewDiary(w, n )
	pc, err := NewPersonalConn() //抽象化されたConfigをインターフェースで作る？
	if err != nil {
		return &usecase.CreateDiaryResponse{
			連携先URL: err.連携先URL
		}
	}
	// d.DiaryGateway.AutoGenerate(sampleDiary)
	return &usecase.CreateDiaryResponse{
		Id: "hoge",
	}, nil
}

func (d *Diary) ConnectNotion(req *struct{}) (res *struct{}, err error) {

	return nil, nil
}
