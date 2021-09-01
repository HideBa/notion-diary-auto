package notion

import (
	"fmt"

	"github.com/HideBa/notion-diary-auto/domain"
	"github.com/HideBa/notion-diary-auto/gateway"
)

type NotionDiary struct {
	notion *Notion
}

func NewNotionDiary() gateway.Diary {
	return &NotionDiary{}
}

func (n *NotionDiary) AutoGenerate(d *domain.Diary) string {
	// ここでNotionにリクエストをなげていく
	fmt.Print("auto generate-------")
	return "hoge"
}
