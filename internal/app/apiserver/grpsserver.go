package apiserver

import (
	"context"
	proto "github.com/Traliaa/http-rest-api/internal/app/proto"
)

type GRPCServer struct {
}

func (s *GRPCServer) UserCreate(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	return &proto.LoginResponse{Email: req.Email, Id: req.Password}, nil
}
