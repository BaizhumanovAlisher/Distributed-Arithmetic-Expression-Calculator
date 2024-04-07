package grpc_client

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"internal/helpers"
	authclient "internal/protos/gen/go/auth_service_v1"
)

const UserId = "userId"

type AuthService struct {
	client authclient.AuthClient
	Secret string
}

func NewAuthService(path string, secret string) (*AuthService, error) {
	if secret == "" {
		return nil, errors.New("secret cannot be empty")
	}

	conn, err := grpc.Dial(path, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := authclient.NewAuthClient(conn)

	return &AuthService{client: client, Secret: secret}, nil
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
				if st.Message() == helpers.InvalidArgumentUserNameErr.Error() {
					return 0, helpers.InvalidArgumentUserNameErr
				}

				if st.Message() == helpers.InvalidArgumentPasswordErr.Error() {
					return 0, helpers.InvalidArgumentPasswordErr
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
				if st.Message() == helpers.InvalidArgumentUserNameErr.Error() {
					return "", helpers.InvalidArgumentUserNameErr
				}

				if st.Message() == helpers.InvalidArgumentPasswordErr.Error() {
					return "", helpers.InvalidArgumentPasswordErr
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
