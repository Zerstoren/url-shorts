package controller

import (
	"github.com/gofiber/fiber/v2"
	featureUser "url-shorts.com/internal/features/User"
	"url-shorts.com/internal/system"
	"url-shorts.com/internal/templates"
)

func setupAuth(app *fiber.App) {
	app.Get("/sign-in", handle(templates.Layout, LoginPage))
	app.Post("/sign-in", handle(templates.Layout, LoginPage))
	app.Get("/sign-up", handle(templates.Layout, RegisterPage))
	app.Post("/sign-up", handle(templates.Layout, RegisterPage))
}

func LoginPage(c *fiber.Ctx) (system.Response, system.ErrorData) {
	errorText := ""
	session, err := store.Get(c)
	if err != nil {
		return nil, system.NewErrorResponse(500, err)
	}

	if session.Get("user") != nil {
		return system.NewResponseRedirect("/"), nil
	}

	if c.Method() == "POST" {
		authError := featureUser.AuthenticateUserByCredentials(
			session,
			c.FormValue("email"),
			c.FormValue("password"),
		)

		if authError == nil {
			return system.NewResponseRedirect("/"), nil
		}

		errorText = *authError
	}

	return featureUser.LoginPage(c, errorText)
}

func RegisterPage(c *fiber.Ctx) (system.Response, system.ErrorData) {
	errorText := ""

	if c.Method() == "POST" {
		session, err := store.Get(c)
		if err != nil {
			return nil, system.NewErrorResponse(500, err)
		}

		registerError := featureUser.RegisterNewUser(
			session,
			c.FormValue("email"),
			c.FormValue("password"),
		)

		if registerError == nil {
			return system.NewResponseRedirect("/"), nil
		}

		errorText = *registerError
	}

	return featureUser.RegisterPage(c, errorText)
}
