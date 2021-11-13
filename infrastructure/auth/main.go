package auth

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var Client *Auth

type FirebaseConfig struct {
	ProjectID string `envconfig:"PROJECT_ID`
}

type Auth struct {
	client *auth.Client
}

type Authenticator interface {
	VerifyToken(idToken string) (uid string, err error)
}

func NewAuthClient(c *FirebaseConfig) Authenticator {
	opt := option.WithCredentialsFile("/credential.json")
	config := &firebase.Config{ProjectID: "dation-c480b"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("ERR: initializing app :%v\n", err)
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatal("ERR: fail to get auth client")
	}
	Client = &Auth{client: client}
	return Client
}

func (a *Auth) VerifyToken(idToken string) (uid string, err error) {
	token, err := a.client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		log.Print("ERR: fail to verify token")
		return "", err
	}
	return token.UID, err
}

type UserInfo struct {
	DisplayName string
}

func (a *Auth) GetUserInfo(ctx context.Context, uid string) *UserInfo {
	u, err := a.client.GetUser(ctx, uid)
	if err != nil {
		log.Print("ERR: fail to get user data")
		return nil
	}
	return &UserInfo{
		DisplayName: u.DisplayName,
	}
}
