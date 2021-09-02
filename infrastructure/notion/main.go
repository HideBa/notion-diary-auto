package notion

type Notion struct {
	authToken  string
	targetDBID string
}

type NotionConfig struct {
	BaseUrl string `envconfig:"BASE_URL"`
	Version string
	Secret  string
}

// func NewNotion(a string, i string) *Notion {
// 	return &Notion{
// 		authToken:  a,
// 		targetDBID: i,
// 	}
// }
