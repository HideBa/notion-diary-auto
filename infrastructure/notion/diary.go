package notion

import (
	"bytes"
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
	jsonBytes, err := ioutil.ReadFile("infrastructure/notion/body.json")
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/pages", n.config.BaseUrl), bytes.NewBuffer(jsonBytes))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", n.config.Secret))
	req.Header.Set("Notion-Version", n.config.Version)
	client := new(http.Client)
	fmt.Print("-----req----", req.Body)
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

type NotionRequest struct {
	parent     NotionParent     `json:"parent"`
	properties []NotionProperty `json:"properties"`
	children   []NotionBlock    `json:"children"`
}

type NotionParent struct {
	parent struct {
		databaseId string `json:"database_id"`
	}
}

type NotionProperty struct {
	name struct {
		title []struct {
			text struct {
				content fmt.Stringer
			} `json:"text"`
		} `json:"title"`
	} `json:"Name"`
}

type NotionChildren []NotionBlock

type NotionBlock struct {
	object    string `json:"object"`
	blockType string `json:"type"`
}
