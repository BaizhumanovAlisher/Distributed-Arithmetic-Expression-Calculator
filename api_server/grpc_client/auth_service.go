package grpc_client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"internal/helpers"
	authclient "protos/gen/go"
)

type AuthService struct {
	client authclient.AuthClient
}

func NewAuthService(path string) (*AuthService, error) {
	conn, err := grpc.Dial(path, grpc.EmptyDialOption{})
	if err != nil {
		return nil, err
	}

	client := authclient.NewAuthClient(conn)

	return &AuthService{client: client}, nil
}

func (a AuthService) Register(ctx context.Context, name, password string) (int64, error) {
	registerResponse, err := a.client.Register(ctx, &authclient.RegisterRequest{Username: name, Password: password})

	if err != nil {
		st, ok := status.FromError(err)

		if ok {
			switch st.Code() {
			case codes.AlreadyExists:
				return 0, helpers.UsernameExistErr
			case codes.Internal:
				return 0, helpers.InternalErr
			case codes.InvalidArgument:
				if st.Message() == helpers.InvalidArgumentUserName.Error() {
					return 0, helpers.InvalidArgumentUserName
				}

				if st.Message() == helpers.InvalidArgumentPassword.Error() {
					return 0, helpers.InvalidArgumentPassword
				}

				return 0, fmt.Errorf(st.Message())
			default:
				return 0, fmt.Errorf(st.Message())
			}
		}

		return 0, err
	}

	return registerResponse.GetUserId(), nil
}

func (a AuthService) Login(ctx context.Context, name, password string) (string, error) {
	loginResponse, err := a.client.Login(ctx, &authclient.LoginRequest{Username: name, Password: password})

	if err != nil {
		st, ok := status.FromError(err)

		if ok {
			switch st.Code() {
			case codes.Unauthenticated:
				return "", helpers.InvalidCredentialsErr
			case codes.Internal:
				return "", helpers.InternalErr
			case codes.NotFound:
				return "", helpers.InvalidCredentialsErr
			case codes.InvalidArgument:
				if st.Message() == helpers.InvalidArgumentUserName.Error() {
					return "", helpers.InvalidArgumentUserName
				}

				if st.Message() == helpers.InvalidArgumentPassword.Error() {
					return "", helpers.InvalidArgumentPassword
				}

				return "", fmt.Errorf(st.Message())
			default:
				return "", fmt.Errorf(st.Message())
			}
		}

		return "", err
	}

	return loginResponse.GetToken(), nil
}
