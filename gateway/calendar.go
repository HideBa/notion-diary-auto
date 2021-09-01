package gateway

import (
	"github.com/HideBa/notion-diary-auto/domain"
	"google.golang.org/genproto/googleapis/type/date"
)

type Calendar interface {
	TodaysCalendar(date.Date) []domain.Calendar
}
