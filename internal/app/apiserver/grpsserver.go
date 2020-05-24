package apiserver

import (
	"context"
	api "github.com/Traliaa/http-rest-api/internal/app/proto"
)

type GRPCServer struct {
}

func (s *GRPCServer) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	return &api.LoginResponse{Email: req.Email, Id: req.Password}, nil
}
