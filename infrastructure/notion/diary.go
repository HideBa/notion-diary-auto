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

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string //TODO: change here for production use
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}
	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrive token from web: %v", err)
	}
	return tok
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

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
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
