package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port              string
	DatabaseURL       string
	JWTSecret         string
	JWTAccessExpiry   int
	JWTRefreshExpiry  int
	CORSOrigin        string
	LogLevel          string
}

func Load() *Config {
	return &Config{
		Port:             getEnv("PORT", "8080"),
		DatabaseURL:      getEnv("DATABASE_URL", "postgres://dangddt:dangddt@localhost:5432/tada_english?sslmode=disable"),
		JWTSecret:        getEnv("JWT_SECRET", "dev-secret-change-in-production"),
		JWTAccessExpiry:  getEnvInt("JWT_ACCESS_EXPIRY", 3600),
		JWTRefreshExpiry: getEnvInt("JWT_REFRESH_EXPIRY", 604800),
		CORSOrigin:       getEnv("CORS_ORIGIN", "*"),
		LogLevel:         getEnv("LOG_LEVEL", "info"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if val := os.Getenv(key); val != "" {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return fallback
}
