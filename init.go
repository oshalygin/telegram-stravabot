package main

import (
	"fmt"
	"net/http"

	"log"

	"github.com/oshalygin/telegram-stravabot/utilities"
	"gopkg.in/telegram-bot-api.v4"
)

var bot *tgbotapi.BotAPI

func initBot(debug bool) {
	configuration := utilities.GetConfiguration()
	bot, err := tgbotapi.NewBotAPI(configuration.TelegramBotAPIKey)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = debug
	log.Printf("Authorized Account: %s", bot.Self.UserName)

}

// InitWebhookConnection initializes a connection to the Telegram BotAPI and polls
// on an interval

func InitWebhookConnection() *tgbotapi.BotAPI {
	configuration := utilities.GetConfiguration()
	const debug = false

	initBot(debug)

	response, err := bot.SetWebhook(tgbotapi.NewWebhook(configuration.WebhookDomain + bot.Token))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)

	webhookInfo, _ := bot.GetWebhookInfo()
	log.Println(webhookInfo)

	updates := bot.ListenForWebhook("/" + bot.Token)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		message := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		message.ReplyToMessageID = update.Message.MessageID

		bot.Send(message)

	}

	return bot
}

// InitPollingConnection initializes a connection to the Telegram BotAPI and polls
// on an interval
func InitPollingConnnection() *tgbotapi.BotAPI {
	const debug = true
	const timeout = 60

	initBot(debug)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeout

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}

	return bot
}

// HandleHealthCheck is a handler which responds back with ok at the healthcheck route
// Healthcheck route: /_ah/health
func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

// HandleRoot is a generic handler which responds back with a welcome page
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Strava Bot!")
}
