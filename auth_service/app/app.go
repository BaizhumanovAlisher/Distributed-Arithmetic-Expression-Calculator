package app

import (
	"auth_service/grpc_server"
	"internal/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCSvr *grpc.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	grpcApp := grpc.New(log, grpcPort, grpc_server.Register)

	return &App{
		GRPCSvr: grpcApp,
	}
}
