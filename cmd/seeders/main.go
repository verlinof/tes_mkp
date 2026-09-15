package main

import (
	"log"
	"time"

	"github.com/verlinof/fiber-project-structure/configs/app_config"
	"github.com/verlinof/fiber-project-structure/configs/db_config"
	"github.com/verlinof/fiber-project-structure/db"
	auth_model "github.com/verlinof/fiber-project-structure/internal/modules/auth/model"
	showtime_model "github.com/verlinof/fiber-project-structure/internal/modules/showtime/model"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	log.Println("Starting database seeding...")

	// 1. Load Configurations
	app_config.Config = app_config.LoadConfig()
	db_config.Config = db_config.LoadConfig()

	// 2. Initialize Database Connection
	db.Init()
	database := db.GetDB()

	// 3. Seed Users (Admin & Customer)
	log.Println("Seeding Users...")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	users := []auth_model.User{
		{
			Name:      "Admin",
			Email:     "admin@example.com",
			Password:  string(hashedPassword),
			Role:      "admin",
			CreatedAt: time.Now(),
		},
		{
			Name:      "User",
			Email:     "user@example.com",
			Password:  string(hashedPassword),
			Role:      "customer",
			CreatedAt: time.Now(),
		},
	}

	for _, u := range users {
		var existing auth_model.User
		if err := database.Where("email = ?", u.Email).First(&existing).Error; err != nil {
			if err := database.Create(&u).Error; err != nil {
				log.Printf("Failed to seed user %s: %v", u.Email, err)
			} else {
				log.Printf("-> User created: %s (%s)", u.Email, u.Role)
			}
		} else {
			log.Printf("-> User %s already exists, skipping", u.Email)
		}
	}

	// 4. Seed Movies
	log.Println("Seeding Movies...")
	movies := []showtime_model.Movie{
		{
			Title:           "Inception",
			DurationMinutes: 148,
			Genre:           "Sci-Fi, Action",
			PosterURL:       "https://image.tmdb.org/t/p/w500/edv5CZvWj09upOsy2Y6IwDhK8bt.jpg",
			CreatedAt:       time.Now(),
		},
		{
			Title:           "Interstellar",
			DurationMinutes: 169,
			Genre:           "Sci-Fi, Adventure, Drama",
			PosterURL:       "https://image.tmdb.org/t/p/w500/gEU2QniE6E77NI6lCU6MxlNBvIx.jpg",
			CreatedAt:       time.Now(),
		},
		{
			Title:           "The Dark Knight",
			DurationMinutes: 152,
			Genre:           "Action, Crime, Drama",
			PosterURL:       "https://image.tmdb.org/t/p/w500/qJ2tW6WMUDux911r6m7haRef0WH.jpg",
			CreatedAt:       time.Now(),
		},
		{
			Title:           "Dune: Part Two",
			DurationMinutes: 166,
			Genre:           "Sci-Fi, Adventure",
			PosterURL:       "https://image.tmdb.org/t/p/w500/1pdfLvkbY9ohJlCjQH2CZjjYVvJ.jpg",
			CreatedAt:       time.Now(),
		},
		{
			Title:           "Agak Laen",
			DurationMinutes: 119,
			Genre:           "Comedy, Horror",
			PosterURL:       "https://image.tmdb.org/t/p/w500/agak_laen_poster.jpg",
			CreatedAt:       time.Now(),
		},
	}

	for _, m := range movies {
		var existing showtime_model.Movie
		if err := database.Where("title = ?", m.Title).First(&existing).Error; err != nil {
			if err := database.Create(&m).Error; err != nil {
				log.Printf("Failed to seed movie %s: %v", m.Title, err)
			} else {
				log.Printf("-> Movie created: %s (ID: %d)", m.Title, m.ID)
			}
		} else {
			log.Printf("-> Movie %s already exists, skipping", m.Title)
		}
	}

	// 5. Seed Cinemas & Studios
	log.Println("Seeding Cinemas and Studios...")
	cinemaData := []struct {
		Cinema  showtime_model.Cinema
		Studios []string
	}{
		{
			Cinema: showtime_model.Cinema{
				Name:      "XXI Grand Indonesia",
				City:      "Jakarta",
				Address:   "Grand Indonesia Shopping Town, Jl. M.H. Thamrin No.1, Jakarta Pusat",
				CreatedAt: time.Now(),
			},
			Studios: []string{"Studio 1", "Studio 2", "IMAX", "Premiere"},
		},
		{
			Cinema: showtime_model.Cinema{
				Name:      "Cinepolis Paragon Mall",
				City:      "Semarang",
				Address:   "Pollux Mall Paragon, Jl. Pemuda No.118, Sekayu, Kota Semarang",
				CreatedAt: time.Now(),
			},
			Studios: []string{"Studio 1", "Studio 2", "Macro XE"},
		},
		{
			Cinema: showtime_model.Cinema{
				Name:      "XXI Tunjungan Plaza",
				City:      "Surabaya",
				Address:   "Tunjungan Plaza 5, Jl. Embong Malang No.1-30, Surabaya",
				CreatedAt: time.Now(),
			},
			Studios: []string{"Studio 1", "Studio 2", "IMAX"},
		},
		{
			Cinema: showtime_model.Cinema{
				Name:      "XXI Beachwalk Bali",
				City:      "Denpasar",
				Address:   "Beachwalk Shopping Center, Jl. Pantai Kuta, Badung, Bali",
				CreatedAt: time.Now(),
			},
			Studios: []string{"Studio 1", "Premiere"},
		},
	}

	for _, item := range cinemaData {
		var cinema showtime_model.Cinema
		if err := database.Where("name = ? AND city = ?", item.Cinema.Name, item.Cinema.City).First(&cinema).Error; err != nil {
			cinema = item.Cinema
			if err := database.Create(&cinema).Error; err != nil {
				log.Printf("Failed to seed cinema %s: %v", item.Cinema.Name, err)
				continue
			}
			log.Printf("-> Cinema created: %s (City: %s, ID: %d)", cinema.Name, cinema.City, cinema.ID)
		} else {
			log.Printf("-> Cinema %s already exists", cinema.Name)
		}

		// Seed Studios for this Cinema
		for _, studioName := range item.Studios {
			var existingStudio showtime_model.Studio
			if err := database.Where("cinema_id = ? AND name = ?", cinema.ID, studioName).First(&existingStudio).Error; err != nil {
				totalSeats := 50
				if studioName == "IMAX" || studioName == "Macro XE" {
					totalSeats = 100
				} else if studioName == "Premiere" {
					totalSeats = 30
				}

				newStudio := showtime_model.Studio{
					CinemaID:   cinema.ID,
					Name:       studioName,
					TotalSeats: totalSeats,
					CreatedAt:  time.Now(),
				}
				if err := database.Create(&newStudio).Error; err != nil {
					log.Printf("Failed to seed studio %s for cinema %d: %v", studioName, cinema.ID, err)
				} else {
					log.Printf("   └── Studio created: %s (Seats: %d, ID: %d)", newStudio.Name, newStudio.TotalSeats, newStudio.ID)
				}
			}
		}
	}

	log.Println("Seeding completed successfully!")
}
