package weather

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/HideBa/notion-diary-auto/domain"
	"github.com/HideBa/notion-diary-auto/gateway"
)

type YahooWeatherConfig struct {
	BaseUrl  string `envconfig:"BASE_URL"`
	ClientId string `envconfig:"CLIENT_ID"`
}
type YahooWeather struct {
	config *YahooWeatherConfig
}

var exampleLat = 35.60904
var exampleLng = 139.648021

const expectedResponseFormat string = "json"

func NewYahooWeather(c *YahooWeatherConfig) gateway.Weather {
	return &YahooWeather{
		config: c,
	}
}

func (y *YahooWeather) TodaysWeather(t time.Time, l domain.Location) (*domain.Weather, error) {
	formattedDate := t.Format("20060102030405")
	fmt.Print("--------", formattedDate)
	req, err := http.NewRequest("GET", y.config.BaseUrl, nil)
	params := req.URL.Query()
	params.Add("appid", y.config.ClientId)
	params.Add("output", expectedResponseFormat)
	params.Add("date", formattedDate)
	params.Add("past", "2")
	params.Add("coordinates", fmt.Sprintf("%s,%s", strconv.FormatFloat(l.Lat, 'f', -1, 64), strconv.FormatFloat(l.Lng, 'f', -1, 64)))
	req.URL.RawQuery = params.Encode()
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Print("------", string(body))
	return nil, nil
}

type WeatherReqest struct {
	appID string
}
