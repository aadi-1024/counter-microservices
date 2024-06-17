package main

import (
	"auth/database"
	"auth/server"
	"authproto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		log.Fatalln(err.Error())
	}
	srv := grpc.NewServer()	

	db, err := database.New(5 * time.Second)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("auth: connected to database")
	authServ := server.AuthServer{}
	authServ.Db = db

	authproto.RegisterAuthRPCServer(srv, authServ)
	log.Println("auth: starting service")
	if err := srv.Serve(lis); err != nil {
		log.Fatalln(err.Error())
	}
}