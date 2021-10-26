package gateway

import "github.com/HideBa/notion-diary-auto/domain"

type Diary interface {
	Create(*domain.Diary) (id string)
}
