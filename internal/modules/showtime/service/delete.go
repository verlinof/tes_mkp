package showtime_service

import (
	"context"
	"errors"
	"fmt"

	showtime_model "github.com/verlinof/fiber-project-structure/internal/modules/showtime/model"
	"gorm.io/gorm"
)

func (s *ShowtimeService) Delete(ctx context.Context, id int64) error {
	var showtime showtime_model.Showtime
	if err := s.db.WithContext(ctx).First(&showtime, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("showtime with id %d not found", id)
		}
		return err
	}

	if err := s.db.WithContext(ctx).Delete(&showtime).Error; err != nil {
		return err
	}

	return nil
}
