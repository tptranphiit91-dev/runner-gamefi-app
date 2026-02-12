package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// DatabaseType represents the type of database to use
type DatabaseType string

const (
	PostgresDB DatabaseType = "postgres"
	MongoDB    DatabaseType = "mongodb"
)

// Config holds application configuration
type Config struct {
	Server       ServerConfig
	Database     DatabaseConfig
	DatabaseType DatabaseType
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string
	Host string
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	// PostgreSQL specific
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string

	// MongoDB specific
	MongoURI     string
	MongoDBName  string
	MongoTimeout int
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if exists
	_ = godotenv.Load()

	dbType := DatabaseType(getEnv("DB_TYPE", "postgres"))

	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Host: getEnv("SERVER_HOST", "0.0.0.0"),
		},
		DatabaseType: dbType,
		Database: DatabaseConfig{
			// PostgreSQL config
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "booking_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),

			// MongoDB config
			MongoURI:     getEnv("MONGO_URI", "mongodb://localhost:27017"),
			MongoDBName:  getEnv("MONGO_DB_NAME", "booking_db"),
			MongoTimeout: getEnvAsInt("MONGO_TIMEOUT", 10),
		},
	}, nil
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt gets an environment variable as int or returns a default value
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// getEnvAsBool gets an environment variable as bool or returns a default value
func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}
