// Just trying to figure out the slack go package github.com/slack-go/slack
package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

const channel = "using-the-slack-api"

func main() {
	key := os.Getenv("SLACK_API_TOKEN")

	if key == "" {
		fmt.Println("SLACK_API_TOKEN environment variable not set")
	}

	api := slack.New(key)

	_, _, err := api.PostMessage(channel, slack.MsgOptionText("This is the message", false))
	if err != nil {
		fmt.Printf("Failed to post: %v", err)
	}
}
