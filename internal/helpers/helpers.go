package helpers

import (
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {
	return slog.New(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelDebug}))
}
