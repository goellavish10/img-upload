package app

import (
	"fmt"
	"os"

	"github.com/goellavish10/img-upload-go/config"
	"github.com/goellavish10/img-upload-go/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupAppAndRun() error {
	// Loading env
	err := config.LoadEnv()

	if err != nil {
		fmt.Println("Error loading env: ", err)
		return err
	}

	// Fiber App Initiation
	app := fiber.New(fiber.Config{
		Prefork:           false,
		StrictRouting:     false,
		CaseSensitive:     true,
		StreamRequestBody: true,
		AppName:           "Image Uploader",
	})

	app.Use(cors.New())

	// Setting up a custom logger
	app.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Kolkata",
	}))

	// Setting up router for application
	router.SetupRouter(app)

	// Swagger Route
	config.SwaggerRoute(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	app.Listen(":" + port)

	return nil

}
