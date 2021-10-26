package notion

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/HideBa/notion-diary-auto/domain"
	"golang.org/x/oauth2"
)

func (n *Notion) PostDiary(d *domain.Diary) string {
	client := getClient(&n.config.OAuth)
	jsonBytes, err := ioutil.ReadFile("infrastructure/notion/body.json")
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/pages", n.config.Base.BaseUrl), bytes.NewBuffer(jsonBytes))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", n.config.Base.Secret))
	req.Header.Set("Notion-Version", n.config.Base.Version)
	// client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Print(err.Error())
	}
	defer res.Body.Close()
	byteArray, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(byteArray))
	// ここでNotionにリクエストをなげていく
	return "hoge"
}

func getClient(config *oauth2.Config) *http.Client {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		panic("ERR: no token") //FIXME: change here
	}
	return config.Client(context.Background(), tok)
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
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
