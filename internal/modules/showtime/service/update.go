package showtime_service

import (
	"context"
	"errors"
	"fmt"
	"time"

	showtime_model "github.com/verlinof/fiber-project-structure/internal/modules/showtime/model"
	"gorm.io/gorm"
)

func (s *ShowtimeService) Update(ctx context.Context, id int64, req showtime_model.UpdateShowtimeRequest) (showtime_model.ShowtimeResponse, error) {
	var showtime showtime_model.Showtime
	if err := s.db.WithContext(ctx).First(&showtime, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return showtime_model.ShowtimeResponse{}, fmt.Errorf("showtime with id %d not found", id)
		}
		return showtime_model.ShowtimeResponse{}, err
	}

	if req.MovieID != nil {
		var movie showtime_model.Movie
		if err := s.db.WithContext(ctx).First(&movie, "id = ?", *req.MovieID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return showtime_model.ShowtimeResponse{}, fmt.Errorf("movie with id %d not found", *req.MovieID)
			}
			return showtime_model.ShowtimeResponse{}, err
		}
		showtime.MovieID = *req.MovieID
	}

	if req.StudioID != nil {
		var studio showtime_model.Studio
		if err := s.db.WithContext(ctx).First(&studio, "id = ?", *req.StudioID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return showtime_model.ShowtimeResponse{}, fmt.Errorf("studio with id %d not found", *req.StudioID)
			}
			return showtime_model.ShowtimeResponse{}, err
		}
		showtime.StudioID = *req.StudioID
	}

	startTime := showtime.StartTime
	if req.StartTime != nil {
		startTime = *req.StartTime
		showtime.StartTime = startTime
	}

	endTime := showtime.EndTime
	if req.EndTime != nil {
		endTime = *req.EndTime
		showtime.EndTime = endTime
	}

	if !endTime.After(startTime) {
		return showtime_model.ShowtimeResponse{}, fmt.Errorf("end_time must be after start_time")
	}

	if req.Price != nil {
		showtime.Price = *req.Price
	}

	if req.Status != nil {
		showtime.Status = *req.Status
	}

	// Check overlap collision
	var count int64
	overlapQuery := s.db.WithContext(ctx).Model(&showtime_model.Showtime{}).
		Where("studio_id = ? AND id != ? AND status != 'cancelled'", showtime.StudioID, id).
		Where("start_time < ? AND end_time > ?", endTime, startTime)
	if err := overlapQuery.Count(&count).Error; err != nil {
		return showtime_model.ShowtimeResponse{}, err
	}
	if count > 0 {
		return showtime_model.ShowtimeResponse{}, fmt.Errorf("studio is already scheduled for another showtime during this time slot")
	}

	showtime.UpdatedAt = time.Now()
	if err := s.db.WithContext(ctx).Save(&showtime).Error; err != nil {
		return showtime_model.ShowtimeResponse{}, err
	}

	// Preload associations for response
	if err := s.db.WithContext(ctx).
		Preload("Movie").
		Preload("Studio.Cinema").
		First(&showtime, "id = ?", id).Error; err != nil {
		return showtime_model.ShowtimeResponse{}, err
	}

	return showtime_model.FormatShowtimeResponse(showtime), nil
}
