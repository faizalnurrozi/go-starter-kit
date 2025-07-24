package middleware

import (
	"fmt"
	"strings"

	"github.com/faizalnurrozi/go-starter-kit/internal/config"
	"github.com/faizalnurrozi/go-starter-kit/internal/errors"
	"github.com/faizalnurrozi/go-starter-kit/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.SendError(c, errors.NewUnauthorizedError())
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		if tokenString == "" {
			return utils.SendError(c, errors.NewUnauthorizedError())
		}

		cfg := config.Get()
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWT.Secret), nil
		})

		if err != nil || !token.Valid {
			fmt.Println(err)
			return utils.SendError(c, errors.NewUnauthorizedError())
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Locals("user_id", claims["user_id"])
			c.Locals("email", claims["email"])
		}

		return c.Next()
	}
}
