package nats

import (
	"context"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"log/slog"
	"strconv"
)

type App struct {
	log *slog.Logger
	js  jetstream.JetStream
}

func NewNATSApp(log *slog.Logger) *App {
	return &App{
		log: log,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "nats.App.Run"

	a.log.With(
		slog.String("op", op),
	)

	connect, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		return err
	}

	js, errConnect := jetstream.New(connect)

	if errConnect != nil {
		return errConnect
	}

	_, errStream := js.CreateStream(context.Background(), jetstream.StreamConfig{
		Name:     "Coins",
		Subjects: []string{"coin.*"},
	})

	if errStream != nil {
		return errStream
	}

	a.js = js

	a.log.Info("NATS is connected!", slog.String("url", a.js.Conn().ConnectedUrl()))

	return nil
}

func (a *App) PublishListingCoin(ctx context.Context, id int) error {
	const op = "nats.App.PublishListingCoin"

	a.log.With(
		slog.String("op", op),
		slog.Int("coinId", id),
	)

	_, err := a.js.Publish(ctx, "coin.listed", []byte(strconv.Itoa(id)))

	if err != nil {
		a.log.Error("failed to publish coin")
	}

	return err
}
