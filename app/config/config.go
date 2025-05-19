package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Username    string
	Password    string
	Database    string
	Host        string
	Dialect     string
	SSLRequired bool
	SSLVerify   bool
}

type ServiceConfig struct {
	BaseURL string
}

type AppConfig struct {
	AppName string
	Env     string
	Port    string

	DB struct {
		Development DatabaseConfig
		Test        DatabaseConfig
		Production  DatabaseConfig
	}

	Services struct {
		Inquiry ServiceConfig
	}
}

func LoadConfig() *AppConfig {
	cfg := &AppConfig{
		AppName: getOrDefault("SERVICE_NAME", "andrian-universe"),
		Env:     getOrDefault("NODE_ENV", "development"),
		Port:    getOrDefault("PORT", "3000"),
	}

	cfg.DB.Development = DatabaseConfig{
		Username: os.Getenv("PGUSER"),
		Password: os.Getenv("PGPASSWORD"),
		Database: os.Getenv("PGDATABASE"),
		Host:     os.Getenv("PGHOST"),
		Dialect:  "postgres",
	}

	cfg.DB.Test = DatabaseConfig{
		Username:    os.Getenv("PGUSER"),
		Password:    os.Getenv("PGPASSWORD"),
		Database:    os.Getenv("PGDATABASE"),
		Host:        os.Getenv("PGHOST"),
		Dialect:     getOrDefault("DB_DIALECT", "postgres"),
		SSLRequired: true,
		SSLVerify:   false,
	}

	cfg.DB.Production = cfg.DB.Test

	// cfg.Services.Inquiry = ServiceConfig{
	// 	BaseURL: os.Getenv("INQUIRY_SERVICE_URL"),
	// }

	return cfg
}

func getOrDefault(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func (cfg *AppConfig) GetActiveDBConfig() DatabaseConfig {
	switch cfg.Env {
	case "test":
		return cfg.DB.Test
	case "production":
		return cfg.DB.Production
	default:
		return cfg.DB.Development
	}
}
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}
