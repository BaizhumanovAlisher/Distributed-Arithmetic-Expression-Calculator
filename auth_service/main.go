package main

import (
	"auth_service/app"
	"internal/helpers"
)

func main() {
	cfg := helpers.MustLoadConfig()

	log := helpers.NewLogger()
	log.Info("starting application")

	application := app.New(log, cfg.GRPCConfig.Port, cfg.Storage.StoragePath, cfg.AuthService.TokenTTL)

	application.GRPCSvr.MustRun()
}
