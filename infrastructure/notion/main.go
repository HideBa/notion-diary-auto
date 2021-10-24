package notion

import (
	"github.com/HideBa/notion-diary-auto/domain"
	"github.com/HideBa/notion-diary-auto/gateway"
)

type NotionPageConfig struct {
	authToken  string
	targetDBID string
}

type NotionConfig struct {
	BaseUrl      string `envconfig:"BASE_URL"`
	Version      string
	Secret       string
	ClientID     string `envconfig:"CLIENT_ID"`
	ClientSecret string `envconfig:"CLIENT_SECRET"`
	Endpoint     string `envconfig:"OAUTH_ENDPOINT"`
}

type NotionDiary struct {
	notion *NotionPageConfig
	config *NotionConfig
}

type PersonalNotionConfig struct {
	accessToken string
}

func NewNotionDiary(c *NotionConfig) gateway.Diary {
	return &NotionDiary{
		config: c,
	}
}

func (n *NotionDiary) AutoGenerate(diary *domain.Diary) (id string) {
	n.CreatePage(diary)
	return "hoge"
}
