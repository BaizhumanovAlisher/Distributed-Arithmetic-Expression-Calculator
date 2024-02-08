package main

import (
	"fmt"
	"log"
	"log/slog"
	"orchestrator/config"
	"orchestrator/storage"
	"os"
	"time"
)

func main() {
	cfg := config.MustLoad()

	logger := setupLogger()

	repo, err := storage.NewSQLite(cfg)
	if err != nil {
		log.Fatal(err)
	}
	//todo: init router
	//todo: run server

	operations, err := repo.ReadAllOperations()
	if err != nil {
		logger.Debug("%s", err)
	}

	for i := 0; i < len(operations); i++ {
		fmt.Println(operations[i])
	}

	time.Sleep(time.Second * 15)
}

func setupLogger() *slog.Logger {
	return slog.New(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelDebug}))
}
