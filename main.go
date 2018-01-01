package main

import (
	"log"

	"fmt"

	"net/http"

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

	response, err := bot.SetWebhook(tgbotapi.NewWebhook("https://www.groklet.com/" + bot.Token))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)

	info, _ := bot.GetWebhookInfo()
	fmt.Println(info)
	http.HandleFunc("/", HandleRoot)
	http.HandleFunc("/_ah/health", HandleHealthCheck)
	updates := bot.ListenForWebhook("/" + bot.Token)

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
