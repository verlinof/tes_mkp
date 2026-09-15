package showtime_model

import "time"

type Cinema struct {
	ID        int       `json:"id" gorm:"column:id;primaryKey"`
	Name      string    `json:"name" gorm:"column:name"`
	City      string    `json:"city" gorm:"column:city"`
	Address   string    `json:"address" gorm:"column:address"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

func (Cinema) TableName() string {
	return "cinemas"
}

type Studio struct {
	ID         int       `json:"id" gorm:"column:id;primaryKey"`
	CinemaID   int       `json:"cinema_id" gorm:"column:cinema_id"`
	Cinema     *Cinema   `json:"cinema,omitempty" gorm:"foreignKey:CinemaID;references:ID"`
	Name       string    `json:"name" gorm:"column:name"`
	TotalSeats int       `json:"total_seats" gorm:"column:total_seats"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
}

func (Studio) TableName() string {
	return "studios"
}

type Movie struct {
	ID              int       `json:"id" gorm:"column:id;primaryKey"`
	Title           string    `json:"title" gorm:"column:title"`
	DurationMinutes int       `json:"duration_minutes" gorm:"column:duration_minutes"`
	Genre           string    `json:"genre" gorm:"column:genre"`
	PosterURL       string    `json:"poster_url" gorm:"column:poster_url"`
	CreatedAt       time.Time `json:"created_at" gorm:"column:created_at"`
}

func (Movie) TableName() string {
	return "movies"
}

type Showtime struct {
	ID        int64     `json:"id" gorm:"column:id;primaryKey"`
	MovieID   int       `json:"movie_id" gorm:"column:movie_id"`
	Movie     *Movie    `json:"movie,omitempty" gorm:"foreignKey:MovieID;references:ID"`
	StudioID  int       `json:"studio_id" gorm:"column:studio_id"`
	Studio    *Studio   `json:"studio,omitempty" gorm:"foreignKey:StudioID;references:ID"`
	StartTime time.Time `json:"start_time" gorm:"column:start_time"`
	EndTime   time.Time `json:"end_time" gorm:"column:end_time"`
	Price     float64   `json:"price" gorm:"column:price"`
	Status    string    `json:"status" gorm:"column:status"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Showtime) TableName() string {
	return "showtimes"
}
