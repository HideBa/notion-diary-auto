package notion

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/HideBa/notion-diary-auto/domain"
	"github.com/HideBa/notion-diary-auto/gateway"
)

type NotionDiary struct {
	notion *Notion
	config *NotionConfig
}

func NewNotionDiary(c *NotionConfig) gateway.Diary {
	return &NotionDiary{
		config: c,
	}
}

func (n *NotionDiary) AutoGenerate(d *domain.Diary) string {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users", n.config.BaseUrl), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", n.config.Secret))
	req.Header.Set("Notion-Version", n.config.Version)
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Print(err.Error())
	}
	defer res.Body.Close()
	byteArray, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(byteArray))
	// ここでNotionにリクエストをなげていく
	fmt.Print("auto generate-------")
	return "hoge"
}
