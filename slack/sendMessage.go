package slack

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

// client for interacting with the Slack API
var client *slack.Client

// specific channel in slack to send messages to
var channel string

func init() {
	// read in the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// grab values from the .env file
	token := os.Getenv("SLACK_API_TOKEN")
	if token == "" {
		log.Fatal("SLACK_API_TOKEN is not set")
	}
	notificationChanel := os.Getenv("CHANNEL")
	if notificationChanel == "" {
		log.Fatal("CHANNEL is not set")
	}

	// use the values or save for later
	client = slack.New(token)
	channel = notificationChanel
}

// SendMessageToSlack sends the provided message to a Slack channel
func SendMessageToSlack(message string) error {
	_, _, err := client.PostMessage(channel, slack.MsgOptionText(message, false))
	if err != nil {
		return fmt.Errorf("failed to post: %v", err)
	}

	return nil
}
