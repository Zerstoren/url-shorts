package featureCreateLink

import (
	featureLink "url-shorts.com/internal/features/Link"
	"url-shorts.com/internal/system"
)

func HandlerCreate(url string) (system.Response, system.ErrorData) {
	settings := system.GetSettings()

	link, err := featureLink.CreateNewLink(url)

	if err != nil {
		return nil, system.NewErrorResponse(503, err)
	}

	content := detailedShortBlock(
		shortResult(
			link.GetTarget(),
			settings.Domain+link.GetShortUrl(),
		),
	)

	return system.NewResponseData(content), nil
}
