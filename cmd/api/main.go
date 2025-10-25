package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/RajaSunrise/pajakku/config"
	"github.com/RajaSunrise/pajakku/internal/databases"
	"github.com/RajaSunrise/pajakku/internal/databases/migrations"
	"github.com/RajaSunrise/pajakku/internal/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize database
	databases.Init()

	// Initialize Redis
	databases.InitRedis()

	// Initialize Migrate
	migrations.Migrate()

	app := fiber.New()

	// Setup Routers
	routers.Routes(app)

	go func() {
		port := config.AppConfig.Server.Port
		if err := app.Listen(":" + port); err != nil {
			log.Fatalf("Error Running Server: %v", err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals
	fmt.Println("Shutting down server")
}
