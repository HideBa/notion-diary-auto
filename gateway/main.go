package gateway

type Gateway struct {
	Diary    *Diary
	Weather  *Weather
	News     *News
	Calendar *Calendar
}

func NewGateway(d *Diary, w *Weather, n *News, c *Calendar) *Gateway {
	return &Gateway{
		Diary:   d,
		Weather: w,
	}
}
