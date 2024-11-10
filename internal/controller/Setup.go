package controller

import (
	"errors"
	"fmt"
	"github.com/a-h/templ"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/session"
	"os"
	"time"
	featureUser "url-shorts.com/internal/features/User"
	"url-shorts.com/internal/system"
)

import "github.com/gofiber/fiber/v2"
import "github.com/gofiber/storage/postgres/v3"

var store *session.Store

func Setup(app *fiber.App) {
	storePg := postgres.New(postgres.Config{
		Host:       "127.0.0.1",
		Port:       5432,
		Username:   os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASS"),
		Database:   os.Getenv("DB_NAME"),
		Table:      "fiber_session",
		GCInterval: 10,
	})

	store = session.New(session.Config{
		CookieHTTPOnly: true,
		Storage:        storePg,
		// CookieSecure: true, for https
		Expiration: time.Hour * 1,
	})

	app.Use(middlewareAuthUser)

	liveReload(app)

	setupAuth(app)
	setupMain(app)
}

func middlewareAuthUser(ctx *fiber.Ctx) error {
	sess, err := store.Get(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	userId := sess.Get("user")

	if userId == nil {
		ctx.Locals("user", nil)
	} else {
		ctx.Locals("user", featureUser.GetUserById(userId.(uint)))
	}

	return ctx.Next()
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

		if content.GetRedirect() != nil {
			ctx.Status(301).Redirect(*content.GetRedirect())
			return nil
		}

		headers := ctx.GetReqHeaders()

		_, ok := headers["Hx-Request"]

		var layoutContent templ.Component

		if content == nil {
			log.Error("No content for page")
			_ = ctx.
				Status(fiber.StatusInternalServerError)
			return errors.New("no content for page")
		}

		if ok {
			layoutContent = content.GetContent()
			ctx.Set("HX-Retarget", "#body-el")

			if content.GetCacheTime() != nil {
				ctx.Set(
					"Cache-Control",
					fmt.Sprintf("private, max-age=%d", int(content.GetCacheTime().Seconds())),
				)
			}
		} else {
			layoutContent = layout(content)
		}

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
