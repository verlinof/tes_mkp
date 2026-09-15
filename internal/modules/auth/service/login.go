package auth_service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	auth_model "github.com/verlinof/fiber-project-structure/internal/modules/auth/model"
	pkg_jwt "github.com/verlinof/fiber-project-structure/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Login(ctx context.Context, req auth_model.LoginRequest) (auth_model.LoginResponse, error) {
	var user auth_model.User
	if err := s.db.WithContext(ctx).Where("email = ?", req.Email).First(&user).Error; err != nil {
		return auth_model.LoginResponse{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return auth_model.LoginResponse{}, err
	}

	claims := jwt.MapClaims{
		"id_user": user.ID,
		"name":    user.Name,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days expiration
	}

	token, err := pkg_jwt.GenerateJWT(&claims)
	if err != nil {
		return auth_model.LoginResponse{}, err
	}

	return auth_model.LoginResponse{
		Token: token,
		User: &auth_model.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
		},
	}, nil
}
