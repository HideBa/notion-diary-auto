package weather

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/HideBa/notion-diary-auto/domain"
	"github.com/HideBa/notion-diary-auto/gateway"
)

type WeatherConfig struct {
	EndPoint string `envconfig:"END_POINT"`
}
type Weather struct {
	config *WeatherConfig
}

var exampleLat = 35.60904
var exampleLng = 139.648021

const expectedResponseFormat string = "json"

func NewWeather(c *WeatherConfig) gateway.Weather {
	return &Weather{
		config: c,
	}
}

func (w *Weather) TodaysWeather(t time.Time, l domain.Location) (*domain.Weather, error) {
	// call this URL https://weather.tsukumijima.net/
	req, err := http.NewRequest("GET", w.config.EndPoint, nil)
	params := req.URL.Query()
	params.Add("city", "130010")
	req.URL.RawQuery = params.Encode()
	header := req.Header
	header.Add("Content-Type", "application/json")
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	// json, err := json.Marshal(&body)
	fmt.Print("------", string(body))
	return nil, nil
}

type WeatherReqest struct {
	appID string
}
