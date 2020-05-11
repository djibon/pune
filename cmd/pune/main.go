package main

import (
	"fmt"
	"pune/internal/handlers"
	
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main(){
	e := echo.New()

	//Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	e.GET("/", handlers.Hello())
	e.GET("/2", handlers.Hello2())
	e.POST("/subscribe-topic", handlers.SubscribeTopic())
	e.POST("/subscribe-group", handlers.SubscribeGroup())
	e.POST("/send-single", handlers.SendSingleDevice())
	e.POST("/send-topic",  handlers.SendTopic())
	e.POST("/send-group",  handlers.SendGroup())

	
	fmt.Println("Start Server")
	e.Logger.Fatal(e.Start(":1234"))
}
