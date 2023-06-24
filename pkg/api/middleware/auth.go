package middleware

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	handlerError "github.com/thnkrn/go-fiber-clean-arch/pkg/api/handler/error"
)

var ErrToken = errors.New("not found token")
var ErrAuthentication = errors.New("unable to successfully authenticate your request")

type Authentication struct {
}

func NewAuthentication() *Authentication {
	return &Authentication{}
}

func validateToken(tokenString string) error {
	if tokenString == "" {
		return handlerError.NewErrorAuthentication(ErrToken)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if err != nil {
		return handlerError.NewErrorAuthentication(err)
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	} else {
		return handlerError.NewErrorAuthentication(ErrAuthentication)
	}
}

func (a *Authentication) Authentication() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		s := ctx.Get("Authorization")
		token := strings.TrimPrefix(s, "Bearer ")

		if err := validateToken(token); err != nil {
			return err
		}
		return ctx.Next()
	}
}

func LoginHandler(c *fiber.Ctx) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		Issuer:    "test",
	})

	ss, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.SendString(ss)
}
