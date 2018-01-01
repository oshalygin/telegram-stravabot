package main

import (
	"log"

	"net/http"

	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	bot, updates := InitConnection()

	http.HandleFunc("/", HandleRoot)
	http.HandleFunc("/_ah/health", HandleHealthCheck)

	go http.ListenAndServe(":8080", nil)

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
