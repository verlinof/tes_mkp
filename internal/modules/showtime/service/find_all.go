package showtime_service

import (
	"context"
	"math"

	showtime_model "github.com/verlinof/fiber-project-structure/internal/modules/showtime/model"
	pkg_success "github.com/verlinof/fiber-project-structure/pkg/success"
	pkg_utils "github.com/verlinof/fiber-project-structure/pkg/utils"
)

var allowedFilters = map[string]string{
	"id":        "showtimes.id",
	"movieId":   "showtimes.movie_id",
	"studioId":  "showtimes.studio_id",
	"status":    "showtimes.status",
	"price":     "showtimes.price",
	"startTime": "showtimes.start_time",
}

var allowedOrderBy = map[string]string{
	"id":        "showtimes.id",
	"startTime": "showtimes.start_time",
	"price":     "showtimes.price",
	"createdAt": "showtimes.created_at",
}

func (s *ShowtimeService) FindAll(ctx context.Context, page, limit int, filters []pkg_utils.Filter, orderBys []pkg_utils.OrderBy) ([]showtime_model.ShowtimeResponse, *pkg_success.Meta, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	query := s.db.WithContext(ctx).Model(&showtime_model.Showtime{})

	// Apply filtering & ordering
	var err error
	query, err = pkg_utils.ApplyFiltersAndOrder(query, filters, orderBys, allowedFilters, allowedOrderBy)
	if err != nil {
		return nil, nil, err
	}

	var totalData int64
	if err := query.Count(&totalData).Error; err != nil {
		return nil, nil, err
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limit)))
	offset := (page - 1) * limit

	var showtimes []showtime_model.Showtime
	if err := query.
		Preload("Movie").
		Preload("Studio.Cinema").
		Offset(offset).
		Limit(limit).
		Find(&showtimes).Error; err != nil {
		return nil, nil, err
	}

	var responses []showtime_model.ShowtimeResponse
	for _, st := range showtimes {
		responses = append(responses, showtime_model.FormatShowtimeResponse(st))
	}

	meta := pkg_success.NewMetaData(page, totalPage, limit, int(totalData))
	return responses, meta, nil
}
