package auth_http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	auth_model "github.com/verlinof/fiber-project-structure/internal/modules/auth/model"
	pkg_error "github.com/verlinof/fiber-project-structure/pkg/error"
	pkg_success "github.com/verlinof/fiber-project-structure/pkg/success"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (h AuthHandler) Register(ctx *fiber.Ctx) error {
	var req auth_model.RegisterRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	if err := h.xValidator.Validate(req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	userResponse, err := h.authService.Register(ctx.Context(), req)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) || err.Error() == "ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)" {
			return ctx.Status(http.StatusConflict).JSON(pkg_error.NewBadRequest(fmt.Errorf("email already registered")))
		}
		return ctx.Status(http.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	return ctx.Status(http.StatusCreated).JSON(pkg_success.SuccessCreateData(userResponse))
}

func (h AuthHandler) Login(ctx *fiber.Ctx) error {
	var req auth_model.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	if err := h.xValidator.Validate(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	loginResponse, err := h.authService.Login(ctx.Context(), req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(pkg_error.NewNotFound(fmt.Errorf("email has not been registered")))
		}
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return ctx.Status(fiber.StatusUnauthorized).JSON(pkg_error.NewUnauthorized(fmt.Errorf("email or password is incorrect")))
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	return ctx.Status(fiber.StatusOK).JSON(pkg_success.SuccessGetData(loginResponse))
}

func (h AuthHandler) GetProfile(ctx *fiber.Ctx) error {
	user, ok := ctx.Locals("token").(*jwt.Token)
	if !ok || user == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(pkg_error.NewUnauthorized(fmt.Errorf("unauthorized")))
	}

	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(pkg_error.NewUnauthorized(fmt.Errorf("invalid token claims")))
	}

	userIDFloat, ok := claims["id_user"].(float64)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(pkg_error.NewUnauthorized(fmt.Errorf("invalid user id in token claims")))
	}

	userResponse, err := h.authService.GetProfile(ctx.Context(), int64(userIDFloat))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(http.StatusNotFound).JSON(pkg_error.NewNotFound(err))
		}
		return ctx.Status(http.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	return ctx.Status(http.StatusOK).JSON(pkg_success.SuccessGetData(userResponse))
}
