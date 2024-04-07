package main

import (
	"expression_solver/app"
	"expression_solver/app/agent_components"
	"google.golang.org/grpc"
	grpc2 "internal/grpc"
	"internal/helpers"
	"internal/protos/gen/go/expression_solver_v1"
	"internal/storage/postgresql"
	logStandard "log"
	"log/slog"
)

func main() {
	cfg := helpers.MustLoadConfig()

	log := helpers.NewLogger()
	log.Info("starting application")

	repository, err := postgresql.NewPostgresql(cfg)
	if err != nil {
		logStandard.Fatal(err)
	}

	agent := agent_components.NewAgent(cfg.Agent.CountOperation)
	expressionManager, err := app.NewExpressionManager(agent, repository)
	if err != nil {
		logStandard.Fatal(err)
	}

	controller := app.NewGrpcController(agent, expressionManager)
	grpcServer := NewExpressionSolverGRPCServer(log, cfg.AuthService.GrpcPort, controller)

	go grpcServer.MustRun()
	_ = helpers.WaitSignal()

	grpcServer.Stop()
	log.Info("application stopped")
}

func NewExpressionSolverGRPCServer(log *slog.Logger, port int, controller *app.GrpcController) *grpc2.BasicGRPCServer {
	gRPCServer := grpc.NewServer()

	expression_solver_v1.RegisterExpressionSolverServer(gRPCServer, controller)
	grpcApp := grpc2.New(log, port, gRPCServer)

	return grpcApp
}
