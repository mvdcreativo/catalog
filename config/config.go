package config

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Name string
	Port string
}

type DatabaseConfig struct {
	Uri  string
	Name string
}

type BucketConfig struct {
	Name     string
	Region   string
	Key      string
	Secret   string
	Endpoint string
	BaseURL  string
	UseSSL   bool
}

type FileValidationConfig struct {
	MaxSizeMB    int64    `mapstructure:"maxSizeMB"`
	AllowedTypes []string `mapstructure:"allowedTypes"`
}

type UploadConfig struct {
	Images FileValidationConfig `mapstructure:"images"`
	Docs   FileValidationConfig `mapstructure:"docs"`
}

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Bucket   BucketConfig
	Upload   UploadConfig
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	env := flag.String("env", "prod", "Ambiente de ejecución (dev, prod)")
	flag.Parse()

	var logMessage string
	var configPath string
	if *env == "dev" {
		configPath = "config/config.dev.yaml"
		logMessage = "😎 Loading development environment..."
	} else {
		configPath = "config/config.yaml"
		logMessage = "🔥 Loading production environment"
	}

	// Leer e interpolar variables de entorno
	content, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("❌ No se pudo leer %s: %v", configPath, err)
	}

	interpolated := os.ExpandEnv(string(content))

	viper.SetConfigType("yaml")
	viper.AutomaticEnv() // permite override por env vars
	if err := viper.ReadConfig(strings.NewReader(interpolated)); err != nil {
		log.Fatalf("❌ Error leyendo config: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("❌ Error parseando configuración: %v", err)
	}

	log.Print(logMessage)

	return &cfg
}
