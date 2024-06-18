package main

import (
	"authproto"
	"counterproto"
)

type App struct {
	Auth authproto.AuthRPCClient
	Counter counterproto.CounterRPCClient
}
