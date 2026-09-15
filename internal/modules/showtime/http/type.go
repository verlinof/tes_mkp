package showtime_http

import (
	showtime_service "github.com/verlinof/fiber-project-structure/internal/modules/showtime/service"
	pkg_validation "github.com/verlinof/fiber-project-structure/pkg/validation"
)

type ShowtimeHandler struct {
	service    showtime_service.ShowtimeService
	xValidator pkg_validation.XValidator
}

func NewHandler(service showtime_service.ShowtimeService, xValidator pkg_validation.XValidator) ShowtimeHandler {
	return ShowtimeHandler{
		service:    service,
		xValidator: xValidator,
	}
}
