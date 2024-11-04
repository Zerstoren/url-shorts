package featureMainPage

import "url-shorts.com/internal/system"

func HandlerMainPage() (system.Response, system.ErrorData) {
	return system.NewResponseData(templatePage()), nil
}
