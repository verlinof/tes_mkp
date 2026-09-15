package showtime_route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/verlinof/fiber-project-structure/db"
	"github.com/verlinof/fiber-project-structure/internal/middleware"
	showtime_http "github.com/verlinof/fiber-project-structure/internal/modules/showtime/http"
)

func InitRoute(router fiber.Router, handler showtime_http.ShowtimeHandler) {
	// Base showtimes route group with JWT Auth & User existence verification
	showtimes := router.Group("/showtimes")

	showtimes.Get("/", handler.FindAll)
	showtimes.Get("/:id", handler.FindByID)

	showtimesadmin := router.Group("/showtimes", middleware.AuthMiddleware(), middleware.CheckUserExists(db.GetDB()))
	// POST, PUT, DELETE: Khusus role Admin saja
	showtimesadmin.Post("/", middleware.RequireAdmin(), handler.Create)
	showtimesadmin.Put("/:id", middleware.RequireAdmin(), handler.Update)
	showtimesadmin.Delete("/:id", middleware.RequireAdmin(), handler.Delete)
}
