package auth

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

type FirebaseConfig struct {
	ProjectID string `envconfig:"PROJECT_ID`
}

type Authenticator struct {
	Client *auth.Client
}

func NewAuthClient(c *FirebaseConfig) *Authenticator {
	ctx := context.Background()
	opt := option.WithCredentialsFile("/credential.json")
	config := &firebase.Config{ProjectID: "dation-c480b"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("ERR: initializing app :%v\n", err)
	}
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatal("ERR: fail to get auth client")
	}
	return &Authenticator{Client: client}
}

func (a *Authenticator) VerifyToken(ctx context.Context, idToken string) (uid string, isValid bool) {
	token, err := a.Client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Print("ERR: fail to verify token")
		return "", false
	}
	return token.UID, true
}

type UserInfo struct {
	DisplayName string
}

func (a *Authenticator) GetUserInfo(ctx context.Context, uid string) *UserInfo {
	u, err := a.Client.GetUser(ctx, uid)
	if err != nil {
		log.Print("ERR: fail to get user data")
		return nil
	}
	return &UserInfo{
		DisplayName: u.DisplayName,
	}
}
