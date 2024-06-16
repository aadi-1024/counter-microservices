package main

import (
	"authproto"
	"log"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var app *App

func main() {
	e := echo.New()
	SetupRouter(e)
	app = &App{}

	client, err := grpc.NewClient("127.0.0.1:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err.Error())
	}

	authClient := authproto.NewAuthRPCClient(client)
	app.Auth = authClient

	if err := e.Start("127.0.0.1:8080"); err != nil {
		log.Fatalln(err.Error())
	}
}
