package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	auth_model "github.com/verlinof/fiber-project-structure/internal/modules/auth/model"
	pkg_error "github.com/verlinof/fiber-project-structure/pkg/error"
)

// RequireRoles checks whether the authenticated user has one of the specified roles
func RequireRoles(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. Check if user entity exists in Locals (populated by CheckUserExists)
		if user, ok := c.Locals("user").(auth_model.User); ok {
			for _, r := range allowedRoles {
				if user.Role == r {
					return c.Next()
				}
			}
			return c.Status(fiber.StatusForbidden).JSON(pkg_error.NewForbidden(fmt.Errorf("access denied: requires %v role", allowedRoles)))
		}

		// 2. Otherwise extract role from JWT claims
		token, ok := c.Locals("token").(*jwt.Token)
		if !ok || token == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(pkg_error.NewUnauthorized(fmt.Errorf("unauthorized")))
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(pkg_error.NewUnauthorized(fmt.Errorf("invalid token claims")))
		}

		role, ok := claims["role"].(string)
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(pkg_error.NewForbidden(fmt.Errorf("access denied: role not found in token")))
		}

		for _, r := range allowedRoles {
			if role == r {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(pkg_error.NewForbidden(fmt.Errorf("access denied: requires %v role", allowedRoles)))
	}
}

// RequireAdmin is a convenience middleware restricting access strictly to admin role
func RequireAdmin() fiber.Handler {
	return RequireRoles("admin")
}
