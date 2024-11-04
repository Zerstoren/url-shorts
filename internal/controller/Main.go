package controller

import (
	"github.com/gofiber/fiber/v2"
	featureCreateLink "url-shorts.com/internal/features/CreateLink"
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
