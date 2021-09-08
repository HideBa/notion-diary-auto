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

func (y *YahooWeather) TodaysWeather(d time.Time, l domain.Location) (*domain.Weather, error) {
	exampleLocation, err := domain.NewLocation(exampleLat, exampleLng)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", y.config.BaseUrl, nil)
	params := req.URL.Query()
	params.Add("appid", y.config.ClientId)
	params.Add("coordinates", fmt.Sprintf("%s,%s", strconv.FormatFloat(exampleLocation.Lat, 'f', -1, 64), strconv.FormatFloat(exampleLocation.Lng, 'f', -1, 64)))
	fmt.Print("---------", params.Get("coordinates"))
	params.Add("output", expectedResponseFormat)
	// メモ：URLエンコードされるときに,までエンコードされてうまく行ってない疑惑
	req.URL.RawQuery = params.Encode()
	fmt.Print("-------", req.URL.String())
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Print("body-----", string(body))
	return nil, nil
}

type WeatherReqest struct {
	appID string
}
