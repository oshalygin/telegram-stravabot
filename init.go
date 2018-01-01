package main

import (
	"fmt"
	"net/http"

	"log"

	"github.com/oshalygin/telegram-stravabot/utilities"
	"gopkg.in/telegram-bot-api.v4"
)

func initBot(debug bool) *tgbotapi.BotAPI {
	configuration := utilities.GetConfiguration()
	bot, err := tgbotapi.NewBotAPI(configuration.TelegramBotAPIKey)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = debug
	log.Printf("Authorized Account: %s", bot.Self.UserName)

	return bot
}

// InitWebhookConnection initializes a connection to the Telegram BotAPI and polls
// on an interval

func InitWebhookConnection() (*tgbotapi.BotAPI, tgbotapi.UpdatesChannel) {
	configuration := utilities.GetConfiguration()
	const debug = false

	bot := initBot(debug)

	response, err := bot.SetWebhook(tgbotapi.NewWebhook(configuration.WebhookDomain + bot.Token))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)

	webhookInfo, _ := bot.GetWebhookInfo()
	log.Println(webhookInfo)

	updates := bot.ListenForWebhook("/" + bot.Token)

	return bot, updates
}

// InitPollingConnection initializes a connection to the Telegram BotAPI and polls
// on an interval
func InitPollingConnection() (*tgbotapi.BotAPI, tgbotapi.UpdatesChannel) {
	const debug = true
	const timeout = 60

	bot := initBot(debug)

	// Remove the webhook if it exists
	bot.RemoveWebhook()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeout

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		log.Panic(err)
	}

	return bot, updates
}

// InitConnection will initialize the connection and determine if the
// webhook connection should be initialized or the polling connection
// based on the environment variable, environment
func InitConnection() (*tgbotapi.BotAPI, tgbotapi.UpdatesChannel) {
	configuration := utilities.GetConfiguration()
	if configuration.Environment == "production" {
		return InitWebhookConnection()
	} else {
		return InitPollingConnection()
	}
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
