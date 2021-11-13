package app

import (
	"errors"

	"github.com/HideBa/notion-diary-auto/infrastructure/auth"

	"github.com/labstack/echo/v4"
)

const authScheme = "Bearer"

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				c.Error(err)
			}

			idToken, err := tokenFromHeader(c, "Authorization", authScheme)
			if err != nil {
				return err
			}
			_, err = auth.Client.VerifyToken(idToken) //Refactor: not use instance directly
			if err != nil {
				return err
			}
			return nil
		}
	}
}

func tokenFromHeader(c echo.Context, header string, authScheme string) (string, error) {
	auth := c.Request().Header.Get(header)
	l := len(authScheme)
	if len(auth) > l+1 && auth[:l] == authScheme {
		return auth[l+1:], nil
	}

	return "", errors.New("missing token")
}
