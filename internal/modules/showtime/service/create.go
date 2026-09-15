package showtime_service

import (
	"context"
	"errors"
	"fmt"

	showtime_model "github.com/verlinof/fiber-project-structure/internal/modules/showtime/model"
	"gorm.io/gorm"
)

func (s *ShowtimeService) Create(ctx context.Context, req showtime_model.CreateShowtimeRequest) (showtime_model.ShowtimeResponse, error) {
	if !req.EndTime.After(req.StartTime) {
		return showtime_model.ShowtimeResponse{}, fmt.Errorf("end_time must be after start_time")
	}

	// 1. Verify Movie exists
	var movie showtime_model.Movie
	if err := s.db.WithContext(ctx).First(&movie, "id = ?", req.MovieID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return showtime_model.ShowtimeResponse{}, fmt.Errorf("movie with id %d not found", req.MovieID)
		}
		return showtime_model.ShowtimeResponse{}, err
	}

	// 2. Verify Studio exists
	var studio showtime_model.Studio
	if err := s.db.WithContext(ctx).Preload("Cinema").First(&studio, "id = ?", req.StudioID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return showtime_model.ShowtimeResponse{}, fmt.Errorf("studio with id %d not found", req.StudioID)
		}
		return showtime_model.ShowtimeResponse{}, err
	}

	// 3. Check for overlapping showtimes in the same studio
	var count int64
	overlapQuery := s.db.WithContext(ctx).Model(&showtime_model.Showtime{}).
		Where("studio_id = ? AND status != 'cancelled'", req.StudioID).
		Where("start_time < ? AND end_time > ?", req.EndTime, req.StartTime)
	if err := overlapQuery.Count(&count).Error; err != nil {
		return showtime_model.ShowtimeResponse{}, err
	}
	if count > 0 {
		return showtime_model.ShowtimeResponse{}, fmt.Errorf("studio is already scheduled for another showtime during this time slot")
	}

	status := req.Status
	if status == "" {
		status = "scheduled"
	}

	showtime := showtime_model.Showtime{
		MovieID:   req.MovieID,
		StudioID:  req.StudioID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Price:     req.Price,
		Status:    status,
	}

	if err := s.db.WithContext(ctx).Create(&showtime).Error; err != nil {
		return showtime_model.ShowtimeResponse{}, err
	}

	showtime.Movie = &movie
	showtime.Studio = &studio

	return showtime_model.FormatShowtimeResponse(showtime), nil
}
