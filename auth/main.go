package main

import (
	"auth/database"
	"auth/server"
	"authproto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatalln(err.Error())
	}
	srv := grpc.NewServer()	

	db, err := database.New()
	if err != nil {
		log.Fatalln(err.Error())
	}

	authServ := server.AuthServer{}
	authServ.Db = db

	authproto.RegisterAuthRPCServer(srv, server.AuthServer{})

	if err := srv.Serve(lis); err != nil {
		log.Fatalln(err.Error())
	}
}