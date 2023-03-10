package main

import (
	"fmt"
	"time"

	"github.com/hunterheston/honeybadger/slack"
	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes and handlers
	e.POST("/record", handleRecord)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}

/////////////////
// handler logic
/////////////////

type RecordInformationRequest struct {
	RecordType    string    `json:"RecordType"`
	Type          string    `json:"Type"`
	TypeCode      int       `json:"TypeCode"`
	Name          string    `json:"Name"`
	Tag           string    `json:"Tag"`
	MessageStream string    `json:"MessageStream"`
	Description   string    `json:"Description"`
	Email         string    `json:"Email"`
	From          string    `json:"From"`
	BouncedAt     time.Time `json:"BouncedAt"`
}

type RecordInformationResponse struct {
	NotificationSent bool   `json:"NotificationSent"`
	Message          string `json:"Message"`
}

// Handle POST requests and optionally send a message to a Slack channel
func handleRecord(c echo.Context) error {

	// Parse the request body.
	recordInfo := RecordInformationRequest{}
	if err := c.Bind(&recordInfo); err != nil {
		return err
	}

	// default response is not to send a notification
	responseInfo := RecordInformationResponse{
		NotificationSent: false,
		Message:          "Message did not meet criteria for sending a notification.",
	}

	if shouldNotifySlack(recordInfo) {
		err := slack.SendMessageToSlack(fmt.Sprintf("Spam message from %q sent to %q at %v", recordInfo.From, recordInfo.Email, recordInfo.BouncedAt))
		if err != nil {
			// Might want to omit the full error message
			responseInfo.Message = fmt.Sprintf("Failed to send notification to slack: %v", err)
			return c.JSON(500, responseInfo)
		}

		responseInfo.NotificationSent = true
		responseInfo.Message = "Message met criteria for sending a notification."
	}

	return c.JSON(200, responseInfo)
}

// simple
func shouldNotifySlack(recordInfo RecordInformationRequest) bool {
	// Very simple check.
	// Could turn this into a named constant or structured type later.
	// Could also check the values of additional fields.
	return recordInfo.Type == "SpamNotification"
}
