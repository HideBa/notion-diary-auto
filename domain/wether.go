package domain

type Weather struct {
	name string
}

const (
	cloudy = "cloudy"
	snow   = "snow"
	sunny  = "sunny"
	rainny = "rainy"
)

func NewWeather(w string) *Weather {
	var weather Weather
	switch w {
	case cloudy:
		weather.name = cloudy
	case snow:
		weather.name = snow
	case sunny:
		weather.name = sunny
	case rainny:
		weather.name = rainny
	default:
		weather.name = sunny
	}
	return &weather
}

func (w *Weather) Name() string {
	return w.name
}
