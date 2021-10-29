package interactor

import (
	"github.com/HideBa/notion-diary-auto/domain"
	"github.com/HideBa/notion-diary-auto/gateway"
	"github.com/HideBa/notion-diary-auto/usecase"
)

type Interactor struct {
	Diary usecase.IDiary
	User  usecase.IUser
}

func NewInteractor(d *gateway.Diary, w *gateway.Weather, n *gateway.News, c *gateway.Calendar, ur *domain.UserRepository) *Interactor {
	return &Interactor{
		Diary: NewDiary(d, w, n, c),
		User:  NewUser(ur),
	}
}
