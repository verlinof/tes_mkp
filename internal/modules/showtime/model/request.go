package showtime_model

import "time"

type CreateShowtimeRequest struct {
	MovieID   int       `json:"movie_id" validate:"required"`
	StudioID  int       `json:"studio_id" validate:"required"`
	StartTime time.Time `json:"start_time" validate:"required"`
	EndTime   time.Time `json:"end_time" validate:"required"`
	Price     float64   `json:"price" validate:"required,gt=0"`
	Status    string    `json:"status" validate:"omitempty,oneof=scheduled ongoing completed cancelled"`
}

type UpdateShowtimeRequest struct {
	MovieID   *int       `json:"movie_id" validate:"omitempty"`
	StudioID  *int       `json:"studio_id" validate:"omitempty"`
	StartTime *time.Time `json:"start_time" validate:"omitempty"`
	EndTime   *time.Time `json:"end_time" validate:"omitempty"`
	Price     *float64   `json:"price" validate:"omitempty,gt=0"`
	Status    *string    `json:"status" validate:"omitempty,oneof=scheduled ongoing completed cancelled"`
}
