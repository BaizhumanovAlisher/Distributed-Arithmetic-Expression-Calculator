package main

import (
	"internal/grpc"
	"internal/grpc/auth"

	// "internal/app/grpc_app"
	"log/slog"
	"time"
)

type App struct {
	GRPCSvr *grpc.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	grpcApp := grpc.New(log, grpcPort, auth.Register)

	return &App{
		GRPCSvr: grpcApp,
	}
}
