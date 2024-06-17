package main

import (
	"authproto"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var app *App

func main() {
	e := echo.New()
	app = &App{}

	time.Sleep(10 * time.Second)
	client, err := grpc.NewClient("auth:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err.Error())
	}
	if client == nil {
		log.Fatalln("fucky wucky")
	}

	authClient := authproto.NewAuthRPCClient(client)
	app.Auth = authClient

	SetupRouter(e)
	if err := e.Start("0.0.0.0:8080"); err != nil {
		log.Fatalln(err.Error())
	}
}
