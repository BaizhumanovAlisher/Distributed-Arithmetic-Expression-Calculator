package main

import (
	"internal/helpers"
)

func main() {
	cfg := helpers.MustLoadConfig()

	log := helpers.NewLogger()
	log.Info("starting application")

	application := New(log, cfg.GRPCConfig.Port, cfg.Storage.StoragePath, cfg.AuthService.TokenTTL)

	application.GRPCSvr.MustRun()
}
