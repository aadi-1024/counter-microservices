package server

import (
	"context"
	"counter/database"
	cp "counterproto"
)

type CounterRPCServer struct {
	cp.UnimplementedCounterRPCServer
	Database *database.Database
}

func (c CounterRPCServer) GetValue(ctx context.Context, req *cp.Request) (*cp.Response, error) {
	resp := &cp.Response{}

	val, err := c.Database.GetValue(int(req.UserId))
	resp.Value = int32(val)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Message = "successful"
	}

	return resp, err
}

func (c CounterRPCServer) Update(ctx context.Context, req *cp.Request) (*cp.Response, error) {
	resp := &cp.Response{}

	val, err := c.Database.UpdateValue(int(req.UserId), int(req.Value), req.Task)
	resp.Value = int32(val)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Message = "successful"
	}
	return resp, err
}
