package showtime_service

import (
	"gorm.io/gorm"
)

type ShowtimeService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) ShowtimeService {
	return ShowtimeService{
		db: db,
	}
}
