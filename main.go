package main

import (
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes and handlers
	e.POST("/record", handleRecord)

	// Start server
	e.Logger.Fatal(e.Start(":3001"))
}

/////////////////
// handler logic
/////////////////

type RecordInformation struct {
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

// Handle POST requests and optionally send a message to a Slack channel
func handleRecord(c echo.Context) error {
	recordInfo := RecordInformation{}
	if err := c.Bind(&recordInfo); err != nil {
		return err
	}
	return c.JSON(200, recordInfo)
}
