package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

type AppConfig struct {
	DB            DBConfig
	Env           string
	JWT_KEYS_PATH string
}

func LoadConfig() AppConfig {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg := AppConfig{
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "app_db"),
			Port:     getEnv("DB_PORT", "5432"),
		},
		Env:           getEnv("APP_ENV", "developer"),
		JWT_KEYS_PATH: getEnv("JWT_KEYS_PATH", "keys"),
	}
	return cfg
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
