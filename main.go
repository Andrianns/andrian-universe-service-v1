package main

import (
	"log"

	"github.com/Andrianns/andrian-universe-service-v1/app/config"
	"github.com/Andrianns/andrian-universe-service-v1/app/models"
	router "github.com/Andrianns/andrian-universe-service-v1/app/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	appCfg := config.LoadConfig()
	config.InitDB(appCfg)
	config.InitClients(appCfg)

	config.DB.AutoMigrate(&models.User{})

	app := fiber.New()
	router.SetupRoutes(app, appCfg)

	port := appCfg.Port
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
