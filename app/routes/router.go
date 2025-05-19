package router

import (
	"github.com/Andrianns/andrian-universe-service-v1/app/config"
	"github.com/Andrianns/andrian-universe-service-v1/app/controllers"
	"github.com/Andrianns/andrian-universe-service-v1/app/repository"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, cfg *config.AppConfig) {
	userRepo := repository.NewUserRepository()
	userController := controllers.NewUserController(userRepo)

	userGroup := app.Group("/users")
	userGroup.Get("/", userController.GetUsers)
	userGroup.Get("/:id", userController.GetUser)
	userGroup.Post("/", userController.CreateUser)
	userGroup.Put("/:id", userController.UpdateUser)
	userGroup.Delete("/:id", userController.DeleteUser)

	// documentController :=

	// app.Post("/cv", documentController.UploadCV)
	// app.Get("/cv", documentController.GetCV)
}
