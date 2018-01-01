package utilities

import "os"

type Configuration struct {
	TelegramBotAPIKey string
	WebhookDomain     string
}

// GetConfiguration returns back the hydrated configuration file which is based on various environment variables
func GetConfiguration() Configuration {
	c := Configuration{
		TelegramBotAPIKey: os.Getenv("TELEGRAM_API_KEY"),
		WebhookDomain:     "https://www.groklet.com/",
	}

	return c
}
