package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	featureLink "url-shorts.com/internal/features/Link"
	featureMainPage "url-shorts.com/internal/features/MainPage"
	"url-shorts.com/internal/system"
	"url-shorts.com/internal/templates"
)

func setupMain(app *fiber.App) {
	app.Post("/create", handle(templates.Layout, Create))
	app.Get("/:code", Redirect)
	app.Get("/", handle(templates.Layout, Main))
}

func Main(ctx *fiber.Ctx) (system.Response, system.ErrorData) {
	return featureMainPage.HandlerMainPage(ctx)
}

func Create(ctx *fiber.Ctx) (system.Response, system.ErrorData) {
	url := ctx.FormValue("link")
	return featureLink.HandlerCreate(url)
}

func Redirect(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	log.Debugf("Try open page: %s", code)
	link, err := featureLink.GetLinkByCode(code)

	if err != nil {
		return err
	}

	return ctx.Redirect(link)
}
