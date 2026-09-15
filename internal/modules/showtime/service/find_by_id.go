package showtime_service

import (
	"context"
	"errors"
	"fmt"

	showtime_model "github.com/verlinof/fiber-project-structure/internal/modules/showtime/model"
	"gorm.io/gorm"
)

func (s *ShowtimeService) FindByID(ctx context.Context, id int64) (showtime_model.ShowtimeResponse, error) {
	var showtime showtime_model.Showtime
	if err := s.db.WithContext(ctx).
		Preload("Movie").
		Preload("Studio.Cinema").
		First(&showtime, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return showtime_model.ShowtimeResponse{}, fmt.Errorf("showtime with id %d not found", id)
		}
		return showtime_model.ShowtimeResponse{}, err
	}

	return showtime_model.FormatShowtimeResponse(showtime), nil
}
