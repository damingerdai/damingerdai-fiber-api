package server

import (
	"fmt"

	"github.com/dammingerdai/damingerdai-fiber-api/pkg/settings"
	"github.com/gofiber/fiber/v2"

	accessLogger "github.com/gofiber/fiber/v2/middleware/logger"
	accessRecover "github.com/gofiber/fiber/v2/middleware/recover"
)

const DEFAULT_HTTP_PORT = 8080

type server struct {
	httpPort int
	app      *fiber.App
}

func NewServer(s settings.ServerSettingS) *server {
	app := fiber.New()
	if s.RunMode == "debug" {
		app.Use(accessLogger.New())
		app.Use(accessRecover.New())
	}
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	app.All("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).JSON("pong")
	})

	svr := server{
		app:      app,
		httpPort: DEFAULT_HTTP_PORT,
	}
	if s.HttpPort > 0 {
		svr.httpPort = s.HttpPort
	}

	return &svr
}

func (s *server) Run(addrs ...int) error {
	if len(addrs) == 0 {
		return s.app.Listen(s.addrs())
	}
	return s.app.Listen(fmt.Sprintf(":%d", addrs[0]))
}

func (s *server) addrs() string {
	return fmt.Sprintf(":%d", s.httpPort)
}
