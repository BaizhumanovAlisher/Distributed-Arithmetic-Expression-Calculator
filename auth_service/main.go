package main

import (
	"auth_service/app"
	"internal/helpers"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := helpers.MustLoadConfig()

	log := helpers.NewLogger()
	log.Info("starting application")

	application := app.New(log, cfg.GRPCConfig.Port, cfg.Storage.StoragePath, cfg.AuthService.TokenTTL)

	go application.GRPCSvr.MustRun()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCSvr.Stop()
	log.Info("application stopped")
}
