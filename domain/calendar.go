package domain

import (
	"google.golang.org/genproto/googleapis/type/date"
)

type Calendar struct {
	date     date.Date
	planName string
	location string
}

func NewCalendar(d *date.Date, pn string, l string) *Calendar {
	return &Calendar{
		date:     *d,
		planName: pn,
		location: l,
	}
}

func (c *Calendar) Date() date.Date {
	return c.date
}

func (c *Calendar) PlanName() string {
	return c.planName
}

func (c *Calendar) Location() string {
	return c.location
}
