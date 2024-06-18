package main

import (
	"counter/database"
	"counter/server"
	"counterproto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:8082")
	if err != nil {
		log.Panicln(err.Error())
	}

	db, err := database.New(5*time.Second)
	if err != nil {
		log.Panicln(err.Error())
	}

	log.Println("connected to DB")

	srv := grpc.NewServer()
	counterRPC := &server.CounterRPCServer{}
	counterRPC.Database = db
	counterproto.RegisterCounterRPCServer(srv, counterRPC)

	log.Println("starting server")
	if err := srv.Serve(l); err != nil {
		log.Panicln(err.Error())
	}
}
