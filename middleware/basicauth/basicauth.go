package basicauth

import (
	"encoding/base64"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// New creates a new middleware handler
func New(config Config) fiber.Handler {
	// ... (existing config setup)
	return func(c *fiber.Ctx) error {
		auth := c.Get(fiber.HeaderAuthorization)
		const prefix = "Basic "

		if len(auth) < len(prefix) || !strings.EqualFold(auth[:len(prefix)], prefix) {
			return config.Unauthorized(c)
		}

		payload, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
		if err != nil {
			return config.Unauthorized(c)
		}

		pair := string(payload)
		colonIndex := strings.IndexByte(pair, ':')
		if colonIndex == -1 {
			return config.Unauthorized(c)
		}

		username := pair[:colonIndex]
		password := pair[colonIndex+1:]

		// ... (existing logic to validate username/password)
		return c.Next()
	}
}