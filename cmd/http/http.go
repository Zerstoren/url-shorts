package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"url-shorts.com/internal/controller"
)

func main() {
	app := fiber.New(fiber.Config{})

	controller.Setup(app)

	app.Static("/public", "./public")

	if err := app.Listen(":3010"); err != nil {
		log.Errorf("server start error, %v", err.Error())
	}
}
