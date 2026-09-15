package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/verlinof/fiber-project-structure/db"
	auth_http "github.com/verlinof/fiber-project-structure/internal/modules/auth/http"
	auth_route "github.com/verlinof/fiber-project-structure/internal/modules/auth/http/route"
	auth_service "github.com/verlinof/fiber-project-structure/internal/modules/auth/service"
	showtime_http "github.com/verlinof/fiber-project-structure/internal/modules/showtime/http"
	showtime_route "github.com/verlinof/fiber-project-structure/internal/modules/showtime/http/route"
	showtime_service "github.com/verlinof/fiber-project-structure/internal/modules/showtime/service"
	pkg_validation "github.com/verlinof/fiber-project-structure/pkg/validation"
)

func InitRoute(app *fiber.App) {
	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "healthy",
			"service": "api-service",
		})
	})

	v1 := app.Group("/v1")

	// Global Dependencies
	validator := pkg_validation.NewXValidator(db.GetDB())

	// Services
	authService := auth_service.NewService(db.GetDB())
	showtimeService := showtime_service.NewService(db.GetDB())

	// Auth Module Routes
	authHandler := auth_http.NewHandler(authService, validator)
	auth_route.InitRoute(v1, authHandler)

	// Showtime Module Routes
	showtimeHandler := showtime_http.NewHandler(showtimeService, validator)
	showtime_route.InitRoute(v1, showtimeHandler)
}
