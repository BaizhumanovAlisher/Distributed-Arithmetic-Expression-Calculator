package main

import (
	"log/slog"
	"orchestrator/config"
	"os"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger()

	log.Debug("Hello logger")
	log.Debug("%+v", cfg)
	//todo: init storage
	//todo: init router
	//todo: run server
}

func setupLogger() *slog.Logger {
	var log = slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	return log
}
