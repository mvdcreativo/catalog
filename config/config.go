package config

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	AppName      string
	Port         string
	DbUser       string
	DbPassword   string
	DbHost       string
	DbName       string
	DbConnParams string
	DbCluster    string
}

// LoadConfig loads configuration from environment variables or .env file
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		AppName:      getEnv("APP_NAME", ""),
		Port:         getEnv("PORT", "8080"),
		DbUser:       getEnv("USER_DB", "default_user"),
		DbPassword:   getEnv("PASSWORD_DB", "default_pass"),
		DbHost:       getEnv("DB_HOST", "localhost"),
		DbName:       getEnv("DB_NAME", "test_db"),
		DbConnParams: getEnv("DB_CONNECTION_PARAMS", ""),
		DbCluster:    getEnv("DB_CLUSTER", "cluster0"),
	}
}

// GetDBURI construye dinámicamente la URI de conexión a MongoDB
func GetDBURI() string {
	cfg := LoadConfig()
	dbUser := cfg.DbUser
	dbPass := cfg.DbPassword
	dbhost := cfg.DbHost
	dbConnParams := cfg.DbConnParams
	dbCluster := cfg.DbCluster

	// Escapar caracteres especiales en la contraseña para evitar errores en la URI
	escapedPassword := url.QueryEscape(dbPass)

	// Construir la URI segura para MongoDB
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/%sappName=%s",
		dbUser, escapedPassword, dbhost, dbConnParams, dbCluster)

	return uri
}

// getEnv retrieves environment variables or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
