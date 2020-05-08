package handlers

import (
	"context"
	"net/http"
	"log"
	"github.com/labstack/echo/v4"

	"pune/internal/libs"
)

type message struct {
	Status string `json:"status"`
	Message string `json:"message"` 
}
func SubscribeTopic() func(echo.Context) error {
	return func(c echo.Context) error {
		app := libs.InitializeAppWithServiceAccount()
		
		ctx := context.Background()
		client, err := app.Messaging(ctx)
		if err != nil {
			//log.Fatalf("error getting Messaging client: %v\n", err)
			msg := &message{
				Status: "0",
				Message: err.Error(),
			}
			return c.JSON(http.StatusOK, msg)
		}

		topic := "TEST"
		registrationTokens := []string{
			"1","2",
		}
		_, err = client.SubscribeToTopic(ctx, registrationTokens, topic)
		log.Println(client)
		if err != nil {
			//log.Fatalf("error getting Messaging client: %v\n", err)
			msg := &message{
				Status: "0",
				Message: err.Error(),
			}
			return c.JSON(http.StatusOK, msg)
		}

		msg := &message{
			Status: "1",
			Message: "OK",
		}
		return c.JSON(http.StatusOK, msg)
	}
}

