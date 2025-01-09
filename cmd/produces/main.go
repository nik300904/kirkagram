package main

import (
	"kirkagram/internal/config"
	k "kirkagram/internal/kafka"
	"kirkagram/internal/lib/logger/handlers/slogpretty"
	"log/slog"
)

func main() {
	cfg := config.New()

	log := setupLogger(cfg.Env)

	log.Info("Starting application")
	log.Info("Current address", slog.String("port", cfg.HttpServe.Address))

	producer := k.NewProducer(*cfg, log)
	_ = producer
}
