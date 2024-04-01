package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type BasicGRPCServer struct {
	log                  *slog.Logger
	registeredGRPCServer *grpc.Server
	port                 int
}

func New(log *slog.Logger, port int, registeredGRPCServer *grpc.Server) *BasicGRPCServer {
	return &BasicGRPCServer{
		log:                  log,
		registeredGRPCServer: registeredGRPCServer,
		port:                 port,
	}
}

func (a *BasicGRPCServer) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *BasicGRPCServer) Run() error {
	const op = "grpc_app.Run"

	log := a.log.With(
		slog.String("op", op),
		slog.Int("port", a.port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	log.Info("starting gRPC server", slog.String("addr", l.Addr().String()))

	if err := a.registeredGRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	return nil
}

func (a *BasicGRPCServer) Stop() {
	const op = "grpc_app.Stop"

	a.log.With(slog.String("op", op)).Info("stopping gRPC server", slog.Int("port", a.port))

	a.registeredGRPCServer.GracefulStop()
}
