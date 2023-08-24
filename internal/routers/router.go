package routers

import (
	"github.com/gofiber/fiber/v2"

	accessLogger "github.com/gofiber/fiber/v2/middleware/logger"
	accessRecover "github.com/gofiber/fiber/v2/middleware/recover"
)

func NewRouter() *fiber.App {
	app := fiber.New()
	app.Use(accessLogger.New())
	app.Use(accessRecover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	app.All("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).JSON("pong")
	})
	return app
}
