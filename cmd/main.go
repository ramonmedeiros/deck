package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/caarlos0/env"
	"github.com/ramonmedeiros/deck/internal/app/rest"
	"github.com/ramonmedeiros/deck/internal/pkg/deck"
)

type config struct {
	Port string `env:"PORT" envDefault:"8080"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	logger.Enabled(context.Background(), slog.LevelError)

	httpServer := rest.New(cfg.Port, logger, deck.NewManager())
	httpServer.Serve()
}
