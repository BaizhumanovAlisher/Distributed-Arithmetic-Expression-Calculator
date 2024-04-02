package app

import (
	"context"
	"errors"
	"internal/helpers"
	"internal/model"
	"internal/storage/postgresql"
	"log/slog"
	"time"
)

type Auth interface {
	Login(ctx context.Context, name string, password string) (string, error)
	Register(ctx context.Context, name string, password string) (int64, error)
}

type JWTAuth struct {
	log        *slog.Logger
	tokenTTL   time.Duration
	repo       *postgresql.PostgresqlDB
	passHasher *PassHasher
	secret     string
}

func NewJWTAuth(log *slog.Logger, tokenTTL time.Duration, repo *postgresql.PostgresqlDB, hasher *PassHasher, secret string) *JWTAuth {
	return &JWTAuth{
		log:        log,
		tokenTTL:   tokenTTL,
		repo:       repo,
		passHasher: hasher,
		secret:     secret,
	}
}

func (j *JWTAuth) Login(ctx context.Context, name string, password string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (j *JWTAuth) Register(ctx context.Context, name string, password string) (int64, error) {
	hashed, err := j.passHasher.GenerateHash(password)
	if err != nil {
		return 0, err
	}

	user := model.NewUser(name, hashed)

	id, err := j.repo.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, helpers.UsernameExistErr) {
			return 0, helpers.UsernameExistErr
		}

		return 0, err
	}

	return id, nil
}
