package featureMainPage

import (
	"github.com/gofiber/fiber/v2"
	featureUser "url-shorts.com/internal/features/User"
	"url-shorts.com/internal/system"
)

func HandlerMainPage(ctx *fiber.Ctx) (system.Response, system.ErrorData) {
	user := featureUser.GetAuthenticateUser(ctx)
	return system.NewResponseData(templatePage(user)), nil
}
