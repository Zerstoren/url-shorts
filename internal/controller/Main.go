package controller

import (
	"github.com/gofiber/fiber/v2"
	featureCreateLink "url-shorts.com/internal/features/CreateLink"
	featureLink "url-shorts.com/internal/features/Link"
	featureMainPage "url-shorts.com/internal/features/MainPage"
	"url-shorts.com/internal/system"
)

func Main(_ *fiber.Ctx) (system.Response, system.ErrorData) {
	return featureMainPage.HandlerMainPage()
}

func Create(ctx *fiber.Ctx) (system.Response, system.ErrorData) {
	url := ctx.FormValue("link")
	return featureCreateLink.HandlerCreate(url)
}

func Redirect(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	link, err := featureLink.GetLinkByCode(code)

	if err != nil {
		return err
	}

	return ctx.Redirect(link)
}
