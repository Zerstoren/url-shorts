package controller

import (
	"github.com/a-h/templ"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/session"
	"time"
	"url-shorts.com/internal/system"
	"url-shorts.com/internal/templates"
)

import "github.com/gofiber/fiber/v2"

var store *session.Store

func Setup(app *fiber.App) {
	store = session.New(session.Config{
		CookieHTTPOnly: true,
		// CookieSecure: true, for https
		Expiration: time.Hour * 1,
	})

	liveReload(app)

	app.Get("/", handle(templates.Layout, Main))
	app.Post("/create", handle(templates.Layout, Create))
}

func handle(
	layout func(data system.Response) templ.Component,
	method func(ctx *fiber.Ctx) (system.Response, system.ErrorData),
) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		content, err := method(ctx)

		if err != nil {
			log.Error(err.Error())

			_ = ctx.
				Status(err.Code()).
				SendString(err.ErrorText())

			return err.Error()
		}

		layoutContent := layout(content)
		return adaptor.HTTPHandler(templ.Handler(layoutContent))(ctx)
	}
}

func liveReload(app *fiber.App) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				break
			}

			if err = c.WriteMessage(mt, msg); err != nil {
				break
			}
		}
	}))
}
