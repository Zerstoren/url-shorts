package system

import "os"

type Settings struct {
	Domain string
}

var settings *Settings

func GetSettings() *Settings {
	if settings == nil {
		settings = &Settings{
			Domain: os.Getenv("DOMAIN"),
		}
	}

	return settings
}
