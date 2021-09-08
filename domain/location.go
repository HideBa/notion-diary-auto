package domain

import "errors"

type Location struct {
	Lat float64
	Lng float64
}

func NewLocation(lat float64, lng float64) (*Location, error) {
	if !(isLatValid(lat) && isLngValid(lng)) {
		return nil, errors.New("lat or lng values is invalid")
	}
	return &Location{
		Lat: lat,
		Lng: lng,
	}, nil
}

func isLatValid(lat float64) bool {
	return lat > -90 && lat < 90
}

func isLngValid(lng float64) bool {
	return lng > -180 && lng < 180
}
