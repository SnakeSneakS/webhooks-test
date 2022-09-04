package main

import (
	"log"

	"github.com/snakesneaks/webhooks-test/config"
	"github.com/snakesneaks/webhooks-test/discord"
	"github.com/snakesneaks/webhooks-test/slack"
)

func main() {
	config := config.GetConfig()

	if err := slack.Send(config.SLACK_WEBHOOK_URL, "test-message"); err != nil {
		log.Println(err)
	} else {
		log.Println("success: send message to slack by webhook")
	}

	if err := discord.Send(config.DISCORD_WEBHOOK_URL, "test-message"); err != nil {
		log.Println(err)
	} else {
		log.Println("success: send message to discord by webhook")
	}
}
