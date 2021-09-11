package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
)

func main() {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatal("error initializing app: %v", err)
	}
}
