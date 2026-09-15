package showtime_http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	showtime_model "github.com/verlinof/fiber-project-structure/internal/modules/showtime/model"
	pkg_error "github.com/verlinof/fiber-project-structure/pkg/error"
	pkg_success "github.com/verlinof/fiber-project-structure/pkg/success"
	pkg_utils "github.com/verlinof/fiber-project-structure/pkg/utils"
)

func (h ShowtimeHandler) Create(c *fiber.Ctx) error {
	var req showtime_model.CreateShowtimeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	if err := h.xValidator.Validate(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	resp, err := h.service.Create(c.Context(), req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	return c.Status(http.StatusCreated).JSON(pkg_success.SuccessCreateData(resp))
}

func (h ShowtimeHandler) FindAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	filters, orderBys := pkg_utils.ParseQueryParams(c)

	responses, meta, err := h.service.FindAll(c.Context(), page, limit, filters, orderBys)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	return c.Status(http.StatusOK).JSON(pkg_success.SuccessPaginationData(responses, meta))
}

func (h ShowtimeHandler) FindByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(pkg_error.NewBadRequest(fmt.Errorf("invalid id parameter")))
	}

	resp, err := h.service.FindByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(pkg_error.NewNotFound(err))
	}

	return c.Status(http.StatusOK).JSON(pkg_success.SuccessGetData(resp))
}

func (h ShowtimeHandler) Update(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(pkg_error.NewBadRequest(fmt.Errorf("invalid id parameter")))
	}

	var req showtime_model.UpdateShowtimeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	if err := h.xValidator.Validate(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	resp, err := h.service.Update(c.Context(), id, req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	return c.Status(http.StatusOK).JSON(pkg_success.SuccessUpdateData(id, resp))
}

func (h ShowtimeHandler) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(pkg_error.NewBadRequest(fmt.Errorf("invalid id parameter")))
	}

	if err := h.service.Delete(c.Context(), id); err != nil {
		return c.Status(http.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	return c.Status(http.StatusOK).JSON(pkg_success.SuccessDeleteData(id))
}
