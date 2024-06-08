package api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/voltgizerz/POS-restaurant/config"
	"github.com/voltgizerz/POS-restaurant/pkg/logger"
)

func NewServer(cfg config.API) {
	app := fiber.New()

	app.Get("/ping", func(c fiber.Ctx) error {
		return c.SendString("pong")
	})

	logger.LogStdErr.Fatal(app.Listen(":" + cfg.PORT))
}
