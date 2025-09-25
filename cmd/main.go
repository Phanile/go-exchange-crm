package main

import (
	"github.com/Phanile/go-exchange-crm/internal/app"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

// TODO: Postgres saves listings
// TODO: NATS consumes trades from go-exchange-trades

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	application := app.NewApp(log)

	go application.GRPCApp.MustRun()
	go application.NATSApp.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCApp.Stop()
}
