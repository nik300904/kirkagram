package main

import (
	"fmt"
	"kirkagram/internal/config"
	"kirkagram/internal/lib/logger/handlers/slogpretty"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
	envDev   = "dev"
)

func main() {
	// TODO: инициализация конфига
	cfg := config.New()

	log := setupLogger(cfg.Env)

	log.Info("Starting application")
	log.Info("Current address", slog.String("port", cfg.HttpServe.Address))

	fmt.Println("THX!")

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{SlogOpts: &slog.HandlerOptions{Level: slog.LevelDebug}}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
