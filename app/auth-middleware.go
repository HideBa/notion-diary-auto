package app

import (
	"io/ioutil"

	firebaseauth "github.com/mondora/firebase-auth-echo-middleware"

	"github.com/labstack/echo/v4"
)

type AuthConfig struct {
	CredentialFilePath string `envconfig:"CREDENTIAL_FILE_PATH"`
}

// func optionsSkipper(c echo.Context) bool {
// }

func NewFirebaseAuthMiddleware(c *AuthConfig) echo.MiddlewareFunc {
	byte, err := ioutil.ReadFile(c.CredentialFilePath)
	if err != nil {
		panic(err)
	}
	fb := firebaseauth.WithConfig(firebaseauth.Config{
		// Skipper: optionsSkipper,
		CredentialJSON: byte,
	})
	return fb
}
