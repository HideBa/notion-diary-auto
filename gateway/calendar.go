package gateway

import (
	"time"

	"github.com/HideBa/notion-diary-auto/domain"
)

type Calendar interface {
	TodaysCalendar(time.Time) []domain.Calendar
}
