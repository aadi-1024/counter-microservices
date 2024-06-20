package main

import (
	"authproto"
	"counterproto"
	"loggerclient"
)

type App struct {
	Auth authproto.AuthRPCClient
	Counter counterproto.CounterRPCClient
	Logger *loggerclient.Logger
}
