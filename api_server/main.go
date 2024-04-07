package main

import (
	"api_server/app"
	"api_server/grpc_client"
	"internal/helpers"
	"internal/storage"
	"internal/storage/postgresql"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	cfg := helpers.MustLoadConfig()
	logger := helpers.NewLogger()

	repo, err := postgresql.NewPostgresql(cfg)
	if err != nil {
		log.Fatal(err)
	}

	redis, err := storage.NewRedisDB(
		cfg.QuickAccessStorage.Address,
		cfg.QuickAccessStorage.Password,
		cfg.QuickAccessStorage.DB,
		cfg.QuickAccessStorage.TTL,
	)

	if err != nil {
		log.Fatal(err)
	}

	authService, err := grpc_client.NewAuthService(cfg.AuthService.Path, cfg.AuthService.Secret)
	if err != nil {
		log.Fatal(err)
	}

	expressionSolver, err := grpc_client.NewExpressionSolver(cfg.ExpressionSolver.Path)
	if err != nil {
		log.Fatal(err)
	}

	application := app.NewApplication(logger, repo, redis, authService, expressionSolver, repo, repo)
	router := application.Routes()

	logger.Info("start server", slog.String("address", cfg.HTTPServer.Address))

	srv := &http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
