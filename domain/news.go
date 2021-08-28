package domain

import "google.golang.org/genproto/googleapis/type/datetime"

type News struct {
	name     string
	url      string
	datetime datetime.DateTime
}

func NewNews(n string, u string, d *datetime.DateTime) *News {
	return &News{
		name:     n,
		url:      u,
		datetime: *d,
	}
}

func (n *News) Name() string {
	return n.name
}

func (n *News) Url() string {
	return n.url
}

func (n *News) Datetime() datetime.DateTime {
	return n.datetime
}
