package middleware

import (
	"time"

	"github.com/faizalnurrozi/go-starter-kit/internal/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		latency := time.Since(start)

		logger.WithFields(logrus.Fields{
			"method":     c.Method(),
			"path":       c.Path(),
			"status":     c.Response().StatusCode(),
			"latency":    latency.String(),
			"ip":         c.IP(),
			"user_agent": c.Get("User-Agent"),
		}).Info("HTTP Request")

		return err
	}
}
