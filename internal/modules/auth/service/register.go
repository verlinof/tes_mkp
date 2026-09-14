package auth_service

import (
	"context"

	auth_model "github.com/verlinof/fiber-project-structure/internal/modules/auth/model"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Register(ctx context.Context, req auth_model.RegisterRequest) (auth_model.UserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return auth_model.UserResponse{}, err
	}

	user := auth_model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := s.db.WithContext(ctx).Create(&user).Error; err != nil {
		return auth_model.UserResponse{}, err
	}

	return auth_model.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}
