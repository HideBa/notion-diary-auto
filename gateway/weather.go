package gateway

import (
	"time"

	"github.com/HideBa/notion-diary-auto/domain"
)

type Weather interface {
	TodaysWeather(time.Time, domain.Location) (*domain.Weather, error)
}
