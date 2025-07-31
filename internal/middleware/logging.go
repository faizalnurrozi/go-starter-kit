package middleware

import (
	"os"
	"time"

	"github.com/faizalnurrozi/go-starter-kit/internal/logger"
	"github.com/gofiber/fiber/v2"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		latency := time.Since(start)

		logger.InfoStructured("Publish event", "event", map[string]any{
			"event_id":         c.Locals("event_id"), // pastikan kamu inject dulu di ctx
			"ip_source":        c.IP(),
			"actor":            c.Locals("actor"), // bisa diisi dari JWT atau ctx
			"api_name":         c.OriginalURL(),
			"url":              c.OriginalURL(), // atau full URL jika tersedia
			"access_time":      start.Format("2006-01-02 15:04:05"),
			"http_status_code": c.Response().StatusCode(),
			"error_code":       "00",      // hardcoded, sesuaikan dengan err handler
			"error_message":    "Success", // sesuaikan juga
			"elapsed_time":     float64(latency.Microseconds()) / 1000.0,
			"pid":              os.Getpid(),
		})

		/* logger.WithFields(logrus.Fields{
			"method":     c.Method(),
			"path":       c.Path(),
			"status":     c.Response().StatusCode(),
			"latency":    latency.String(),
			"ip":         c.IP(),
			"user_agent": c.Get("User-Agent"),
		}).Info("HTTP Request") */

		return err
	}
}
