package config

import (
	"os"
)

type struct DB {
	Host string
	User string
	Password string
	Name string
	Port string
	SslMode string
}

type struc Config {
	Database DB
}

const Config{
	
}

const (
	// Configurable variables
	AppPort          = getEnv("APP_PORT", "8080")
	DatabaseHost     = getEnv("DB_HOST", "localhost")
	DatabaseUser     = getEnv("DB_USER", "username")
	DatabasePassword = getEnv("DB_PASSWORD", "password")
	DatabaseName     = getEnv("DB_NAME", "local")
	DatabasePort     = getEnv("DB_PORT", "5432")
	DatabaseSslMode  = getEnv("DB_SSL", "disable")
	TimeZone         = getEnv("TZ", "America/Santiago")
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
