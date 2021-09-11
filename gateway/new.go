package gateway

import (
	"time"

	"github.com/HideBa/notion-diary-auto/domain"
)

type News interface {
	TodaysNews(time.Time) ([]domain.News, error)
}
