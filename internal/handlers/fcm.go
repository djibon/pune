package handlers

import (
	"context"
	"net/http"
	"github.com/labstack/echo/v4"
	"fmt"
	
	"pune/internal/libs"
	"firebase.google.com/go/messaging"
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

type Token struct {
	DeviceId string `json:"device_id" validate:"required"`
}

type Topic struct {
	Topic string `json:"topic" validate:"required"`
}

//send to single device.
func SendSingleDevice() func(echo.Context) error {
	return func(c echo.Context) error {
		//TODO APP PUT IN CONTEXT?
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
		token := new(Token)

		if err := c.Bind(token); err != nil {
			msg := &Message{
				Status: "0",
				Message: err.Error(),
			}
			return c.JSON(http.StatusOK, msg)
		}

		registrationToken := token.DeviceId
		
		// See documentation on defining a message payload.
		message := &messaging.Message{
			Data: map[string]string{
				"score": "850",
				"time":  "2:45",
			},
			Token: registrationToken,
		}

		// Send a message to the device corresponding to the provided
		// registration token.
		response, err := client.Send(ctx, message)
		if err != nil {
			msg := &Message{
				Status: "0",
				Message: err.Error(),
			}
			return c.JSON(http.StatusOK, msg)
		}
		
		// Response is a message ID string.
		fmt.Println("Successfully sent message:", response)
		msg := &Message{
			Status: "1",
			Message: "OK",
		}
		
		return c.JSON(http.StatusOK, msg)

	}
}

//send to topic
func SendToTopic() func(echo.Context) error {
	return func(c echo.Context) error {
		//TODO APP PUT IN CONTEXT?
		app := libs.InitializeAppWithServiceAccount()
		ctx := context.Background()
		client, err := app.Messaging(ctx)
		topic := new(Topic)
		if err := c.Bind(topic); err != nil{
			msg := &Message{
				Status: "0",
				Message: err.Error(),
			}
			return c.JSON(http.StatusOK, msg)

		}

		topic_topic := topic.Topic
		message := &messaging.Message{
			Data: map[string]string{
				"score": "850",
				"time":  "2:45",
			},
			Topic: topic_topic,
		}
		
		response, err := client.Send(ctx, message)
		if err != nil {
			msg := &Message{
				Status: "0",
				Message: err.Error(),
			}
			return c.JSON(http.StatusOK, msg)
		}
		
		// Response is a message ID string.
		fmt.Println("Successfully sent message:", response)

		msg := &Message{
			Status: "1",
			Message: "OK",
		}
		
		return c.JSON(http.StatusOK, msg)

	}
}

//send to group
func SendGroup() func(echo.Context) error {
	return func(c echo.Context) error {
		
		msg := &Message{
			Status: "1",
			Message: "OK",
		}
		
		return c.JSON(http.StatusOK, msg)

	}
}


//subscribe to group
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


//subscribe tot topic
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

