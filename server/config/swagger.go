package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SwaggerRoute(app *fiber.App) {
	// Swagger setup on /swagger endpoint
	app.Get("/swagger/*", swagger.HandlerDefault)
}
