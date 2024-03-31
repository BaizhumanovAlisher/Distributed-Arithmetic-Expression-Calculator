package auth

import (
	"context"
	"google.golang.org/grpc"
	authservicev1 "protos/gen/go"
)

type serverGRPC struct {
	// todo: implement
	authservicev1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	authservicev1.RegisterAuthServer(gRPC, &serverGRPC{})
}

func (s *serverGRPC) Login(ctx context.Context, req *authservicev1.LoginRequest) (*authservicev1.LoginResponse, error) {
	panic("not implemented")
}

func (s *serverGRPC) Register(ctx context.Context, request *authservicev1.RegisterRequest) (*authservicev1.RegisterResponse, error) {
	panic("implement me")
}
