package app

import (
	"context"
	"errors"
	"fmt"
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
	const op = "auth.Login"

	log := j.log.With(
		slog.String("op", op),
		slog.String("name", name),
	)

	log.Info("read user from DB")

	user, err := j.repo.ReadUserByName(ctx, name)
	if err != nil {
		if errors.Is(err, helpers.NoRowsErr) {
			log.Info("no such user")
			return "", helpers.NoRowsErr
		}

		log.Error("error reading user")
		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("comparing password and hash")

	samePassword, err := j.passHasher.Compare(user.HashedPassword, password)
	if err != nil {
		log.Error("error comparing password")
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if !samePassword {
		log.Info("passwords don't match")
		return "", helpers.InvalidCredentialsErr
	}

	log.Info("user logged in successfully")

	token, err := NewToken(user, j.tokenTTL, j.secret)
	if err != nil {
		log.Error("err to generate token", slog.String("err", err.Error()))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("token generated")
	return token, nil
}

func (j *JWTAuth) Register(ctx context.Context, name string, password string) (int64, error) {
	const op = "auth.Register"

	log := j.log.With(
		slog.String("op", op),
		slog.String("name", name),
	)

	log.Info("start registration user")

	hashed, err := j.passHasher.GenerateHash(password)
	if err != nil {
		log.Error("err hashing password", slog.String("err", err.Error()))
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("password was hashed")

	user := model.NewUser(name, hashed)

	log.Info("start create user in database")

	id, err := j.repo.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, helpers.UsernameExistErr) {
			log.Error("user name already exists")
			return 0, helpers.UsernameExistErr
		}

		log.Error("failed to create user", slog.String("err", err.Error()))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("created user", slog.Int64("id", id))
	return id, nil
}
