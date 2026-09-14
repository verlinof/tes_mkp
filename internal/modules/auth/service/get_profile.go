package auth_service

import (
	"context"

	auth_model "github.com/verlinof/fiber-project-structure/internal/modules/auth/model"
)

func (s *AuthService) GetProfile(ctx context.Context, userID int64) (auth_model.UserResponse, error) {
	var user auth_model.User
	if err := s.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error; err != nil {
		return auth_model.UserResponse{}, err
	}

	return auth_model.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}
