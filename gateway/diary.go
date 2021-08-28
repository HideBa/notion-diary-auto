package gateway

import "github.com/HideBa/notion-diary-auto/domain"

type Diary interface {
	AutoGenerate(*domain.Diary) (id string)
}
