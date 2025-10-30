package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/RajaSunrise/pajakku/config"
	"github.com/RajaSunrise/pajakku/internal/databases"
	"github.com/RajaSunrise/pajakku/internal/databases/migrations"
	"github.com/RajaSunrise/pajakku/internal/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func run(shutdown chan os.Signal) error {
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
		logrus.Infof("Server starting on port %s", port)
		if err := app.Listen(":" + port); err != nil {
			logrus.Errorf("Error Running Server: %v", err)
		}
	}()

	if shutdown == nil {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
		<-signals
	} else {
		<-shutdown
	}
	logrus.Info("Shutting down server")
	return nil
}

func main() {
	if err := run(nil); err != nil {
		logrus.Fatal(err)
	}
}
