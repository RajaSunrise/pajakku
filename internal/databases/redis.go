package databases

import (
	"context"
	"fmt"

	"github.com/RajaSunrise/pajakku/config"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var RDB *redis.Client

func InitRedis() {
	logrus.Info("Initializing Redis connection")
	cfg := config.AppConfig.Redis
	RDB = redis.NewClient(&redis.Options{
		Addr:       fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:   cfg.Password,
		DB:         cfg.DB,
		Protocol:   2,  // Use RESP2 to avoid auto mode issues
		ClientName: "", // Disable client name to avoid maint notifications
	})
	// Test connection
	_, err := RDB.Ping(context.Background()).Result()
	if err != nil {
		logrus.Fatalf("Failed to connect to Redis: %v", err)
	}
	logrus.Info("Connected to Redis")
}
