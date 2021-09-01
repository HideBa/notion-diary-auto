package gateway

import (
	"github.com/HideBa/notion-diary-auto/domain"
	"google.golang.org/genproto/googleapis/type/date"
)

type Weather interface {
	TodaysWeather(date.Date) *domain.Weather
}
