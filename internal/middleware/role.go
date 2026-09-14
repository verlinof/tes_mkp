package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// PassThrough middleware placeholder if needed
func PassThrough() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}
