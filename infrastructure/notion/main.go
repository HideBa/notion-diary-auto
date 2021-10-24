package notion

import (
	"errors"
	"fmt"

	"github.com/HideBa/notion-diary-auto/domain"
	"github.com/HideBa/notion-diary-auto/gateway"
	"golang.org/x/oauth2"
)

type NotionPageConfig struct {
	authToken  string
	targetDBID string
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
			RedirectURL:  "localhost:8080/api/v1/notion/callback",
			Endpoint: oauth2.Endpoint{
				AuthURL:  fmt.Sprintf("%v/authorize", r.Endpoint),
				TokenURL: fmt.Sprint("%v/token", r.Endpoint),
			},
		},
	}
}

type PersonalNotionConfig struct {
	accessToken string
}

func NewNotionDiary(r *NotionRawConfig) gateway.Diary {
	return &Notion{
		config: NewNotionConfig(r),
	}
}

func (n *Notion) Create(diary *domain.Diary) (id string) {
	n.PostDiary(diary)
	return "hoge"
}

func (n *Notion) Connect() error {
	tok := getTokenFromWeb(&n.config.OAuth)
	if tok == nil {
		return errors.New("No token!")
	}
	return nil
}
