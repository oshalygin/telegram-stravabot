package main

import (
	"log"

	"github.com/oshalygin/telegram-stravabot/utilities"
	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	configuration := utilities.GetConfiguration()

	bot, err := tgbotapi.NewBotAPI(configuration.TelegramBotAPIKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized Account: %s", bot.Self.UserName)

}
