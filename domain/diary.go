package domain

type Diary struct {
	name     string
	author   string
	news     *News
	calendar *Calendar
	weather  *Weather
}

func NewDiary(n string, a string) *Diary {
	return &Diary{
		name:   n,
		author: a,
	}
}

func (d *Diary) Name() string {
	return d.name
}

func (d *Diary) Author() string {
	return d.author
}

func (d *Diary) News() *News {
	return d.news
}

func (d *Diary) Calendar() *Calendar {
	return d.calendar
}

func (d *Diary) Weather() *Weather {
	return d.weather
}
