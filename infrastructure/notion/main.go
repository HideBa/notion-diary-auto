package notion

type Notion struct {
	authToken  string
	targetDBID string
}

func NewNotion(a string, i string) *Notion {
	return &Notion{
		authToken:  a,
		targetDBID: i,
	}
}
