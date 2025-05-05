package binder

import (
	"fmt"
	"profkom/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

const (
	key       = "token"
	claimsKey = "user"
	UserID    = "userID"
)

type Middleware struct {
	secret string
}

func New(secret string) *Middleware {
	return &Middleware{
		secret: secret,
	}
}

func (m *Middleware) Auth(ctx *fiber.Ctx) error {
	token := ctx.Cookies(key)
	if token == "" {
		return fiber.ErrUnauthorized
	}

	claims, err := m.parseJwt(token)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	ctx.Locals(claimsKey, claims)

	return ctx.Next()
}

func (m *Middleware) parseJwt(jwtToken string) (*models.ClaimsJwt, error) {
	claims := &models.ClaimsJwt{}

	token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method: %v", token.Header["alg"])
		}
		return []byte(m.secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
