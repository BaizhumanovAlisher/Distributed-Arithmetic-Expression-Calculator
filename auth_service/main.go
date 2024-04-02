package main

import (
	"auth_service/app"
	"google.golang.org/grpc"
	grpc2 "internal/grpc"
	"internal/helpers"
	"internal/storage/postgresql"
	logStandard "log"
	"log/slog"
	authservicev1 "protos/gen/go"
)

func main() {
	cfg := helpers.MustLoadConfig()

	log := helpers.NewLogger()
	log.Info("starting application")

	repository, err := postgresql.NewPostgresql(cfg)
	if err != nil {
		logStandard.Fatal(err)
	}

	authService := app.NewJWTAuth(log, cfg.AuthService.TokenTTL, repository, app.NewPassHasher(cfg.AuthService.Cost), cfg.AuthService.Secret)

	grpcServer := NewAuthGRPCServer(log, cfg.AuthService.GrpcPort, authService)

	go grpcServer.MustRun()
	_ = helpers.WaitSignal()

	grpcServer.Stop()
	log.Info("application stopped")
}

func NewAuthGRPCServer(log *slog.Logger, port int, authService *app.JWTAuth) *grpc2.BasicGRPCServer {
	gRPCServer := grpc.NewServer()

	authservicev1.RegisterAuthServer(gRPCServer, app.NewGRPCController(authService))
	grpcApp := grpc2.New(log, port, gRPCServer)

	return grpcApp
}
