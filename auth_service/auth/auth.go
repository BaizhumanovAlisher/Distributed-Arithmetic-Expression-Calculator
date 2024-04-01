package auth

import (
	"context"
	"internal/storage/postgresql"
	"log/slog"
	"time"
)

type Auth interface {
	Login(ctx context.Context, user string, password string) (string, error)
	Register(ctx context.Context, user string, password string) (int64, error)
}

type JWTAuth struct {
}

func NewJWTAuth(*slog.Logger, time.Duration, *postgresql.PostgresqlDB) *JWTAuth {
	return &JWTAuth{}
}

func (J *JWTAuth) Login(ctx context.Context, user string, password string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (J *JWTAuth) Register(ctx context.Context, user string, password string) (int64, error) {
	//TODO implement me
	panic("implement me")
}
