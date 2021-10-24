package notion

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/HideBa/notion-diary-auto/domain"
	"golang.org/x/oauth2"
)

type NotionInternal struct{}

func (n *NotionDiary) AutoGenerate(d *domain.Diary) string {
	token := getToken()
	if token == nil {

	}
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

func getToken() {

}

func handleCallback() {

}

func storeToken() {

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

func (n *NotionDiary) CreateWithOAuth(diary domain.Diary) error {
	return nil
}

func (n *NotionDiary) Connect(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()
	conf := &oauth2.Config{
		RedirectURL:  "http://localhost:8080/api/v1/notion/callback",
		ClientID:     n.config.ClientID,
		ClientSecret: n.config.ClientSecret,
		Scopes:       []string{},
		Endpoint: oauth2.Endpoint{
			AuthURL:  fmt.Sprintf("%v/authorize", n.config.Endpoint),
			TokenURL: fmt.Sprintf("%v/token", n.config.Endpoint),
		},
	}
	url := conf.AuthCodeURL("state")

	return nil

}

func Callback() error {
	return nil
}


type  PersonalConn interface{
	Config() interface{}
}

func (p *PersonalConn) AutoGenerate(d *domain.Diary){

}

func (p *PersonalConn) Conn() {

}

type NoTokenErr struct {
	integrationUrl string
}

func NewPersonalConn(c *NotionConfig)(*PersonalConn, *NoTokenErr) {
	token := GetToken()
	if !token {
		// 連携する
		return nil, &{
			連携先URL: hogehoge
		}
	}
	return &PersonalConn{
		accessToken: token
	}
}