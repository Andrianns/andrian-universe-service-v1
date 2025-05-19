package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg *AppConfig) {
	dbCfg := cfg.GetActiveDBConfig()

	sslMode := "disable"
	if dbCfg.SSLRequired {
		sslMode = "require"
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=%s",
		dbCfg.Host, dbCfg.Username, dbCfg.Password, dbCfg.Database, sslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	fmt.Println("✅ Database connected")
}
