package app

import (
	"github.com/Phanile/go-exchange-crm/internal/app/grpc"
	"github.com/Phanile/go-exchange-crm/internal/services/listing"
	"log/slog"
)

type App struct {
	GRPCApp *grpc.App
	log     *slog.Logger
}

func NewApp(log *slog.Logger) *App {
	service := listing.NewListingService(log)
	grpcApp := grpc.NewGRPCApp(log, service, 5715)

	return &App{
		log:     log,
		GRPCApp: grpcApp,
	}
}
