package domain

import (
	"time"
)

type News struct {
	name     string
	text     string
	category Category
	url      string
	time     time.Time
}

type Category string

const (
	BUSINESS      = "business"
	ENTERTAINMENT = "entertainment"
	HEALTH        = "health"
	SCIENTCE      = "scient"
	TECHNOLOGY    = "technology"
	SPORTS        = "sports"
)

func NewNews(n string, u string, t *time.Time) *News {
	return &News{
		name: n,
		url:  u,
		time: *t,
	}
}

func (n *News) Name() string {
	return n.name
}

func (n *News) Url() string {
	return n.url
}

func (n *News) Datetime() time.Time {
	return n.time
}
