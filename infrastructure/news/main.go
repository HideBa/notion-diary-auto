package news

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/HideBa/notion-diary-auto/domain"
	"github.com/HideBa/notion-diary-auto/gateway"
)

type News struct {
	config *NewsConfig
}

type NewsConfig struct {
	EndPoint string `envconfig:"END_POINT"`
	ApiKey   string `envconfig:"API_KEY"`
}

func NewNews(c *NewsConfig) gateway.News {
	return &News{
		config: c,
	}
}

func (n *News) TodaysNews(t time.Time) ([]domain.News, error) {

	req, err := http.NewRequest("GET", n.config.EndPoint, nil)
	params := req.URL.Query()
	// ?country=jp&apiKey=cf000d6fd61c43aba7533d70bef24090&from=2021-09-10&soryBy=popularity&category=business
	params.Add("apiKey", n.config.ApiKey)
	params.Add("country", "jp")
	params.Add("from", t.Format("2006-01-02"))
	params.Add("category", "business")
	req.URL.RawQuery = params.Encode()
	header := req.Header
	header.Add("Content-Type", "application/json")
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	// json, err := json.Marshal(&body)
	fmt.Print("------", string(body))
	return nil, nil
}
