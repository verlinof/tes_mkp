package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/verlinof/fiber-project-structure/configs/app_config"
	"github.com/verlinof/fiber-project-structure/configs/db_config"
	"github.com/verlinof/fiber-project-structure/db"
	"github.com/verlinof/fiber-project-structure/internal/routes"
	pkg_cron "github.com/verlinof/fiber-project-structure/pkg/cron"
	pkg_email "github.com/verlinof/fiber-project-structure/pkg/email"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Init Global Config
	app_config.Config = app_config.LoadConfig()
	db_config.Config = db_config.LoadConfig()

	// Email Worker
	pkg_email.InitEmailWorker(5, 100, 3)

	// Connect Database & Services
	db.Init()

	// Init Fiber Engine
	app := fiber.New()

	// CRON Job
	scheduler := pkg_cron.NewScheduler()
	pkg_cron.RegisterJobs(scheduler, db.GetDB())
	scheduler.Start()

	// Add CORS Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, HEAD, PUT, DELETE, PATCH",
	}))

	// Init Route
	routes.InitRoute(app)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		addr := ":" + app_config.Config.AppPort
		log.Printf("Server is running on port %s", app_config.Config.AppPort)
		if err := app.Listen(addr); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	<-quit
	log.Println("Server is shutting down...")

	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server failed to shutdown gracefully: %v", err)
	}

	cleanup(db.GetDB(), scheduler)
	log.Println("Server shut down gracefully")
}

func cleanup(db *gorm.DB, scheduler *pkg_cron.Scheduler) {
	// Example: Close database connection
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Error getting underlying database: %v", err)
		} else {
			if err := sqlDB.Close(); err != nil {
				log.Printf("Error closing database: %v", err)
			} else {
				log.Println("Database connection closed successfully")
			}
		}
	}

	scheduler.Stop()
	pkg_email.ShutdownEmailWorker()
}
