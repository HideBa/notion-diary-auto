package notion

type Notion interface {
	CreatePage(string) error
	Connect() error
	Callback() error
}
