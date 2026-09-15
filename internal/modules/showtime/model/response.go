package showtime_model

import "time"

type CinemaResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Address string `json:"address"`
}

type StudioResponse struct {
	ID         int             `json:"id"`
	CinemaID   int             `json:"cinema_id"`
	Cinema     *CinemaResponse `json:"cinema,omitempty"`
	Name       string          `json:"name"`
	TotalSeats int             `json:"total_seats"`
}

type MovieResponse struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	DurationMinutes int    `json:"duration_minutes"`
	Genre           string `json:"genre"`
	PosterURL       string `json:"poster_url"`
}

type ShowtimeResponse struct {
	ID        int64           `json:"id"`
	MovieID   int             `json:"movie_id"`
	Movie     *MovieResponse  `json:"movie,omitempty"`
	StudioID  int             `json:"studio_id"`
	Studio    *StudioResponse `json:"studio,omitempty"`
	StartTime time.Time       `json:"start_time"`
	EndTime   time.Time       `json:"end_time"`
	Price     float64         `json:"price"`
	Status    string          `json:"status"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

func FormatShowtimeResponse(st Showtime) ShowtimeResponse {
	resp := ShowtimeResponse{
		ID:        st.ID,
		MovieID:   st.MovieID,
		StudioID:  st.StudioID,
		StartTime: st.StartTime,
		EndTime:   st.EndTime,
		Price:     st.Price,
		Status:    st.Status,
		CreatedAt: st.CreatedAt,
		UpdatedAt: st.UpdatedAt,
	}

	if st.Movie != nil {
		resp.Movie = &MovieResponse{
			ID:              st.Movie.ID,
			Title:           st.Movie.Title,
			DurationMinutes: st.Movie.DurationMinutes,
			Genre:           st.Movie.Genre,
			PosterURL:       st.Movie.PosterURL,
		}
	}

	if st.Studio != nil {
		studioResp := &StudioResponse{
			ID:         st.Studio.ID,
			CinemaID:   st.Studio.CinemaID,
			Name:       st.Studio.Name,
			TotalSeats: st.Studio.TotalSeats,
		}
		if st.Studio.Cinema != nil {
			studioResp.Cinema = &CinemaResponse{
				ID:      st.Studio.Cinema.ID,
				Name:    st.Studio.Cinema.Name,
				City:    st.Studio.Cinema.City,
				Address: st.Studio.Cinema.Address,
			}
		}
		resp.Studio = studioResp
	}

	return resp
}
