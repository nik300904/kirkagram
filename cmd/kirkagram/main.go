package main

import (
	"kirkagram/internal/config"
	"kirkagram/internal/lib/logger/handlers/slogpretty"
	"kirkagram/internal/service"
	"kirkagram/internal/storage"
	"kirkagram/internal/storage/psgr"
	"kirkagram/internal/transport/rest"
	"log/slog"
	"net/http"
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
	db := storage.New(cfg)

	log := setupLogger(cfg.Env)

	log.Info("Starting application")
	log.Info("Current address", slog.String("port", cfg.HttpServe.Address))

	userRepo := psgr.NewUserStorage(db)
	userService := service.NewUserService(log, userRepo)
	handler := rest.NewHandler(log, userService)

	srv := &http.Server{
		Addr:    cfg.HttpServe.Address,
		Handler: handler.InitRouter(),
	}

	log.Info("SERVER STARTED AT", slog.String("address", cfg.HttpServe.Address))

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

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
