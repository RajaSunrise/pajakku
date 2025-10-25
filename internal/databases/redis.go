package databases

import (
	"fmt"
	"log"

	"github.com/RajaSunrise/pajakku/config"
	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func InitRedis() {
	cfg := config.AppConfig.Redis
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	log.Println("Connected to Redis")
}
