package utilities

import "os"

type Configuration struct {
	TelegramBotAPIKey string
}

// GetConfiguration returns back the hydrated configuration file which is based on various environment variables
func GetConfiguration() Configuration {
	c := Configuration{
		TelegramBotAPIKey: os.Getenv("TELEGRAM_API_KEY"),
	}

	return c
}
