package router

import (
	"github.com/goellavish10/img-upload-go/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	app.Get("/health", handlers.HealthCheck)

	// Handle File Upload Route
	app.Post("/file-upload", handlers.FileUpload)
}
