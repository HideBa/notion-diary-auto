package weather

import "errors"

exampleLat := 35.60904
exampleLng := 139.648021

type YahooWeather struct {
	baseUrl string
}

type WeatherReqest struct {
	appID
}