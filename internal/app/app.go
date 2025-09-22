package app

import (
	"github.com/Phanile/go-exchange-crm/internal/app/grpc"
	"github.com/Phanile/go-exchange-crm/internal/app/nats"
	"github.com/Phanile/go-exchange-crm/internal/services/listing"
	"log/slog"
)

type App struct {
	GRPCApp *grpc.App
	NATSApp *nats.App
	log     *slog.Logger
}

func NewApp(log *slog.Logger) *App {
	natsApp := nats.NewNATSApp(log)
	service := listing.NewListingService(log, natsApp)
	grpcApp := grpc.NewGRPCApp(log, service, 5715)

	return &App{
		log:     log,
		GRPCApp: grpcApp,
		NATSApp: natsApp,
	}
}
