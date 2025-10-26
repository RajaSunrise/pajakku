package config

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Log      LogConfig      `mapstructure:"log"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSLMode  string `mapstructure:"ssl_mode"`
	TimeZone string `mapstructure:"time_zone"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type JWTConfig struct {
	Secret     string `mapstructure:"secret"`
	ExpiryHour int    `mapstructure:"expiry_hour"`
}

var AppConfig Config

func LoadConfig() {
	// Set default values
	viper.SetDefault("server.port", "3000")
	viper.SetDefault("server.mode", "development")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "")
	viper.SetDefault("database.name", "lookuy")
	viper.SetDefault("database.ssl_mode", "disable")
	viper.SetDefault("database.time_zone", "Asia/Jakarta")
	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", 6379)
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.db", 0)
	viper.SetDefault("jwt.secret", "your-secret-key")
	viper.SetDefault("jwt.expiry_hour", 24)
	viper.SetDefault("log.level", "info")

	// Read from config file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// Read from environment variables
	viper.AutomaticEnv()

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found, using defaults and environment variables")
		} else {
			log.Fatalf("Error reading config file: %v", err)
		}
	}

	// Unmarshal config
	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	// Override with environment variables if set
	if port := os.Getenv("SERVER_PORT"); port != "" {
		AppConfig.Server.Port = port
	}
	if mode := os.Getenv("SERVER_MODE"); mode != "" {
		AppConfig.Server.Mode = mode
	}
	if host := os.Getenv("DB_HOST"); host != "" {
		AppConfig.Database.Host = host
	}
	if user := os.Getenv("DB_USER"); user != "" {
		AppConfig.Database.User = user
	}
	if password := os.Getenv("DB_PASSWORD"); password != "" {
		AppConfig.Database.Password = password
	}
	if name := os.Getenv("DB_NAME"); name != "" {
		AppConfig.Database.Name = name
	}
	if timezone := os.Getenv("TimeZone"); timezone != "" {
		AppConfig.Database.TimeZone = timezone
	}
	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		AppConfig.JWT.Secret = secret
	}

	log.Println("Config loaded successfully")

	// Setup logrus
	setupLogrus()
}

func setupLogrus() {
	level, err := logrus.ParseLevel(AppConfig.Log.Level)
	if err != nil {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(level)
	}

	if AppConfig.Server.Mode == "development" {
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
}
