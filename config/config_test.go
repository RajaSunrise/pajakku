package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Save original env vars
	originalEnv := map[string]string{
		"SERVER_PORT": os.Getenv("SERVER_PORT"),
		"SERVER_MODE": os.Getenv("SERVER_MODE"),
		"DB_HOST":     os.Getenv("DB_HOST"),
		"DB_USER":     os.Getenv("DB_USER"),
		"DB_PASSWORD": os.Getenv("DB_PASSWORD"),
		"DB_NAME":     os.Getenv("DB_NAME"),
		"TimeZone":    os.Getenv("TimeZone"),
		"JWT_SECRET":  os.Getenv("JWT_SECRET"),
	}

	// Restore env vars after test
	defer func() {
		for key, value := range originalEnv {
			if value != "" {
				os.Setenv(key, value)
			} else {
				os.Unsetenv(key)
			}
		}
	}()

	tests := []struct {
		name             string
		envVars          map[string]string
		expectedServer   ServerConfig
		expectedDatabase DatabaseConfig
		expectedJWT      JWTConfig
	}{
		{
			name:    "default config",
			envVars: map[string]string{},
			expectedServer: ServerConfig{
				Port: "3000",
				Mode: "development",
			},
			expectedDatabase: DatabaseConfig{
				Host:     "localhost",
				Port:     5432,
				User:     "postgres",
				Password: "",
				Name:     "lookuy",
				SSLMode:  "disable",
				TimeZone: "Asia/Jakarta",
			},
			expectedJWT: JWTConfig{
				Secret:     "your-secret-key",
				ExpiryHour: 24,
			},
		},
		{
			name: "custom env vars",
			envVars: map[string]string{
				"SERVER_PORT": "8080",
				"SERVER_MODE": "production",
				"DB_HOST":     "db.example.com",
				"DB_USER":     "testuser",
				"DB_PASSWORD": "testpass",
				"DB_NAME":     "testdb",
				"JWT_SECRET":  "testsecret",
			},
			expectedServer: ServerConfig{
				Port: "8080",
				Mode: "production",
			},
			expectedDatabase: DatabaseConfig{
				Host:     "db.example.com",
				Port:     5432, // Not overridden
				User:     "testuser",
				Password: "testpass",
				Name:     "testdb",
				SSLMode:  "disable",
				// TimeZone: "Asia/Jakarta",
			},
			expectedJWT: JWTConfig{
				Secret:     "testsecret",
				ExpiryHour: 24,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set env vars for test
			for key, value := range tt.envVars {
				os.Setenv(key, value)
			}

			// Reset AppConfig
			AppConfig = Config{}

			// Load config
			LoadConfig()

			// Assert server config
			assert.Equal(t, tt.expectedServer, AppConfig.Server)

			// Assert database config
			assert.Equal(t, tt.expectedDatabase, AppConfig.Database)

			// Assert JWT config
			assert.Equal(t, tt.expectedJWT, AppConfig.JWT)

			// Assert Redis config (defaults)
			expectedRedis := RedisConfig{
				Host:     "localhost",
				Port:     6379,
				Password: "",
				DB:       0,
			}
			assert.Equal(t, expectedRedis, AppConfig.Redis)
		})
	}
}
