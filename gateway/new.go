package gateway

import (
	"github.com/HideBa/notion-diary-auto/domain"
	"google.golang.org/genproto/googleapis/type/date"
)

type News interface {
	TodaysNews(date.Date) []domain.News
}
