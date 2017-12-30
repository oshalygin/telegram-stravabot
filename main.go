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

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		log.Println(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		message := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		message.ReplyToMessageID = update.Message.MessageID

		bot.Send(message)

	}
}
