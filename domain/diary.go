package domain

type Diary struct {
	name    string
	author  string
	news    *News
	calenar *Calendar
	weather *Weather
}

func NewDiary(n string, a string) *Diary {
	return &Diary{
		name:   n,
		author: a,
	}
}
