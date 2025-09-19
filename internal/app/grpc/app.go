package grpc

import (
	"fmt"
	listingGRPC "github.com/Phanile/go-exchange-crm/internal/grpc/listing"
	"github.com/Phanile/go-exchange-crm/internal/services/listing"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewGRPCApp(log *slog.Logger, listingService *listing.ListingService, port int) *App {
	gRPCServer := grpc.NewServer()
	listingGRPC.Register(gRPCServer, listingService)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpc.app.Run"

	a.log.With(
		slog.String("op", op),
		slog.Int("port", a.port),
	)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	a.log.Info("grpc server is running on", slog.String("addr", listener.Addr().String()))

	return a.gRPCServer.Serve(listener)
}

func (a *App) Stop() {
	const op = "grpc.app.Stop"

	a.log.With(
		slog.String("op", op),
	)

	a.log.Info("grpc server is shutting down")

	a.gRPCServer.GracefulStop()
}
