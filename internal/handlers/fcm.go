package handlers

import (
	"context"
	"net/http"
	"github.com/labstack/echo/v4"

	"pune/internal/libs"
)

type Message struct{
	Status string `json:"status"`
	Message string `json:"message"`
}

type SubscribeData struct{
	Topic string `json:"topic" validate:"required"`
	DeviceId string `json:"device_id" validate:"required"`
}

type GroupData struct{
	Group string `json:"group" validate:"required"`
	DeviceId string `json:"device_id" validate:"required"`
}


func SubscribeGroup() func(echo.Context) error {
	return func(c echo.Context) error {
		gt := new(GroupData)

		if err := c.Bind(gt); err != nil {
			msg := &Message{
				Status: "0",
				Message: err.Error(),
			}
			return c.JSON(http.StatusOK, msg)
		}

		if len(gt.Group) == 0 {
			msg := &Message{
				Status: "0",
				Message: "Group Required",
			}
			return c.JSON(http.StatusOK, msg)
		}

		msg := &Message{
			Status: "1",
			Message: "OK",
		}
		
		return c.JSON(http.StatusOK, msg)
	}
}


func SubscribeTopic() func(echo.Context) error {
	return func(c echo.Context) error {
		app := libs.InitializeAppWithServiceAccount()

		
		ctx := context.Background()
		client, err := app.Messaging(ctx)
		if err != nil {
			msg := &Message{
				Status: "0",
				Message: err.Error(),
			}
			return c.JSON(http.StatusOK, msg)
		}

		st := new(SubscribeData)
		if err = c.Bind(st); err != nil {
			msg := &Message{
				Status: "0",
				Message: err.Error(),
			}
			return c.JSON(http.StatusOK, msg)
		}
		
		topic := st.Topic;
		
		registrationTokens := []string{
			st.DeviceId,
		}
		
		_, err = client.SubscribeToTopic(ctx, registrationTokens, topic)

		if err != nil {
			msg := &Message{
				Status: "0",
				Message: err.Error(),
			}
			return c.JSON(http.StatusOK, msg)
		}

		msg := &Message{
			Status: "1",
			Message: "OK",
		}
		return c.JSON(http.StatusOK, msg)
	}
}

