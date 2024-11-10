package featureUser

import (
	"github.com/gofiber/fiber/v2"
	"url-shorts.com/internal/system"
)

func LoginPage(c *fiber.Ctx, errorText string) (system.Response, system.ErrorData) {
	email := ""
	if c.Method() == "POST" {
		email = c.FormValue("email")
	}

	return system.NewResponseData(loginForm(email, errorText)), nil
}

func RegisterPage(c *fiber.Ctx, errorText string) (system.Response, system.ErrorData) {
	return system.NewResponseData(registerForm(errorText)), nil
}
