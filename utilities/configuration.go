package utilities

import "os"

type Configuration struct {
	TelegramBotAPIKey string
	WebhookDomain     string
	Environment       string
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// GetConfiguration returns back the hydrated configuration file which is based on various environment variables
func GetConfiguration() Configuration {
	return Configuration{
		Environment:       getEnv("ENVIRONMENT", "dev"),
		TelegramBotAPIKey: getEnv("TELEGRAM_API_KEY", ""),
		WebhookDomain:     "https://www.groklet.com/",
	}
}
