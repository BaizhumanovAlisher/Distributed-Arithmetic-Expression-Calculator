package app

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"internal/helpers"
	"internal/validators"
	authservicev1 "protos/gen/go"
)

type GrpcController struct {
	// todo: implement
	authservicev1.UnimplementedAuthServer
	Auth *JWTAuth
}

func NewGRPCController(auth *JWTAuth) *GrpcController {
	return &GrpcController{Auth: auth}
}

func (g *GrpcController) Login(ctx context.Context, req *authservicev1.LoginRequest) (*authservicev1.LoginResponse, error) {
	err := g.ValidateCredentials(req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, err
	}

	token, err := g.Auth.Login(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		if errors.Is(err, helpers.InvalidCredentialsErr) {
			return nil, status.Error(codes.Unauthenticated, helpers.InvalidCredentialsErr.Error())
		}

		if errors.Is(err, helpers.NoRowsErr) {
			return nil, status.Error(codes.NotFound, helpers.NoRowsErr.Error())
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &authservicev1.LoginResponse{
		Token: token,
	}, nil
}

func (g *GrpcController) Register(ctx context.Context, req *authservicev1.RegisterRequest) (*authservicev1.RegisterResponse, error) {
	err := g.ValidateCredentials(req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, err
	}

	id, err := g.Auth.Register(ctx, req.GetUsername(), req.GetPassword())

	if err != nil {
		if errors.Is(err, helpers.UsernameExistErr) {
			return nil, status.Error(codes.AlreadyExists, helpers.UsernameExistErr.Error())
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &authservicev1.RegisterResponse{
		UserId: id,
	}, nil
}

func (g *GrpcController) ValidateCredentials(username, password string) error {
	if err := validators.ValidateUsername(username); err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	if err := validators.ValidatePassword(password); err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	return nil
}
