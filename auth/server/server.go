package server

import (
	"auth/database"
	"authproto"
	"context"

	// "google.golang.org/grpc"
)

type AuthServer struct {
	authproto.UnimplementedAuthRPCServer
	Db *database.Database
}

func (a AuthServer) Login(ctx context.Context, in *authproto.LoginRequest) (*authproto.Response, error) {
	resp := &authproto.Response{}

	if err := a.Db.Login(in.Email, in.Password); err != nil {
		resp.Success = false
		return resp, err
	}

	resp.Success = true
	resp.Message = "successful"
	return resp, nil
}

func (a AuthServer) Register(ctx context.Context, in *authproto.RegisterRequest) (*authproto.Response, error) {
	resp := &authproto.Response{}

	if err := a.Db.Register(in.Email, in.Username, in.Password); err != nil {
		resp.Success = false;
		return resp, err
	}

	resp.Success = true
	resp.Message = "successful"
	return resp, nil
}