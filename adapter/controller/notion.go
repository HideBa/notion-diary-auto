package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type OAuthController struct {
	Notion *Notion
}

func NewOAuthController(r *NotionRawConfig) *OAuthController {
	return &OAuthController{
		Notion: NewNotionOAuth(r),
	}
}

type NotionRawConfig struct {
	BaseUrl      string `envconfig:"BASE_URL"`
	Version      string
	Secret       string
	ClientID     string `envconfig:"CLIENT_ID"`
	ClientSecret string `envconfig:"CLIENT_SECRET"`
	Endpoint     string `envconfig:"OAUTH_ENDPOINT"`
}

type Notion struct {
	// notion *NotionPageConfig
	config *NotionConfig
}

type NotionConfig struct {
	Base  BaseConfing
	OAuth oauth2.Config
}

type BaseConfing struct {
	BaseUrl string
	Version string
	Secret  string
}

func NewNotionConfig(r *NotionRawConfig) *NotionConfig {
	return &NotionConfig{
		Base: BaseConfing{
			BaseUrl: r.BaseUrl,
			Version: r.Version,
			Secret:  r.Secret,
		},
		OAuth: oauth2.Config{
			ClientID:     r.ClientID,
			ClientSecret: r.ClientSecret,
			RedirectURL:  "http://localhost:8080/api/v1/notion/callback",
			Endpoint: oauth2.Endpoint{
				AuthURL:  fmt.Sprintf("%v/authorize", r.Endpoint),
				TokenURL: fmt.Sprintf("%v/token", r.Endpoint),
			},
		},
	}
}

func NewNotionOAuth(r *NotionRawConfig) *Notion {
	return &Notion{
		config: NewNotionConfig(r),
	}
}

func (n *Notion) GetAuthCode(c echo.Context) error {
	authURL := n.config.OAuth.AuthCodeURL("state-token")
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	return c.Redirect(http.StatusFound, authURL)
}

func (n *Notion) GetToken(c echo.Context) error {
	authCode := c.QueryParam("code")
	state := c.QueryParam("state")
	fmt.Print("state---", state)
	tok, err := n.config.OAuth.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrive token from web: %v", err)
	}
	fmt.Print(tok)
	saveToken("token.json", tok)
	return c.JSON(200, "token")
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
