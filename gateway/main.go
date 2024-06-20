package main

import (
	"authproto"
	"counterproto"
	"log"
	"loggerclient"
	"time"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var app *App

func main() {
	e := echo.New()
	app = &App{}

	time.Sleep(5 * time.Second)
	authClient, err := grpc.NewClient("auth:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err.Error())
	}

	auth := authproto.NewAuthRPCClient(authClient)
	app.Auth = auth

	counterClient, err := grpc.NewClient("counter:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err.Error())
	}

	counter := counterproto.NewCounterRPCClient(counterClient)
	app.Counter = counter

	logger, err := loggerclient.NewLogger("gateway", "amqp://guest:guest@rabbitmq:5672")
	if err != nil {
		log.Fatalln(err.Error())
	}
	app.Logger = logger

	SetupRouter(e)
	if err := e.Start("0.0.0.0:8080"); err != nil {
		log.Fatalln(err.Error())
	}
}
