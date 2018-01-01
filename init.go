package main

import (
	"fmt"
	"net/http"
)

// InitWebhookConnection initializes a connection to the Telegram BotAPI and polls
// on an interval

func InitWebhookConnection() {

}

// InitPollingConnection initializes a connection to the Telegram BotAPI and polls
// on an interval
func InitPollingConnnection() {

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
