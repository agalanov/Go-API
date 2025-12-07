package config

import (
	"os"
)

type Config struct {
	Database DatabaseConfig
	JWT      JWTConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type JWTConfig struct {
	SecretKey     string
	AccessExpiry  int // minutes
	RefreshExpiry int // days
}

type ServerConfig struct {
	Port         string
	InternalToken string
}

func Load() *Config {
	return &Config{
		Database: DatabaseConfig{
			Host:     getEnv("POSTGRES_HOST", "localhost"),
			Port:     getEnv("POSTGRES_PORT", "5432"),
			User:     getEnv("POSTGRES_USER", "postgres"),
			Password: getEnv("POSTGRES_PASSWORD", ""),
			DBName:   getEnv("POSTGRES_DATABASE_API", "diflow"),
			SSLMode:  getEnv("POSTGRES_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			SecretKey:     getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
			AccessExpiry:  60,  // 60 minutes
			RefreshExpiry: 30, // 30 days
		},
		Server: ServerConfig{
			Port:         getEnv("PORT", "8080"),
			InternalToken: getEnv("INTERNAL_API_TOKEN", ""),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

