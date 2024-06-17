package server

import (
	"auth/database"
	"authproto"
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	// "google.golang.org/grpc"
)

type AuthServer struct {
	authproto.UnimplementedAuthRPCServer
	Db *database.Database
}

type claims struct {
	Uid int
	Expiry time.Time
	jwt.RegisteredClaims
}

func (a AuthServer) Login(ctx context.Context, in *authproto.LoginRequest) (*authproto.Response, error) {
	resp := &authproto.Response{}

	id, err := a.Db.Login(in.Email, in.Password)
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		return resp, err
	}

	clms := &claims{
		Uid: id,
		Expiry: time.Now().Add(24*time.Hour),
	}

	token := jwt.New(jwt.SigningMethodHS512)
	token.Claims = clms
	ss, err := token.SignedString([]byte("HUGE_SECRET"))

	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		return resp, err
	}

	resp.Success = true
	resp.Message = ss
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