//	@title			Kirkagram API
//	@version		1.0
//	@description	This is the API for Kirkagram application
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8082
//	@BasePath	/api

package main

import (
	httpSwagger "github.com/swaggo/http-swagger"
	_ "kirkagram/docs"
	"kirkagram/internal/config"
	k "kirkagram/internal/kafka"
	"kirkagram/internal/lib/logger/handlers/slogpretty"
	"kirkagram/internal/service"
	"kirkagram/internal/storage"
	"kirkagram/internal/storage/psgr"
	S3Storage "kirkagram/internal/storage/s3"
	"kirkagram/internal/transport/rest"
	"kirkagram/internal/transport/rest/handlers"
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
	cfg := config.New()
	db := storage.New(cfg)
	S3Client := storage.NewS3Client()

	log := setupLogger(cfg.Env)

	log.Info("Starting application")
	log.Info("Current address", slog.String("port", cfg.HttpServe.Address))

	userRepo := psgr.NewUserStorage(db)
	postRepo := psgr.NewPostStorage(db)
	likeRepo := psgr.NewLikeStorage(db)
	followRepo := psgr.NewFollowStorage(db)
	s3Repo := S3Storage.NewUserS3Storage(S3Client)
	producer := k.NewProducer(cfg, log)

	userService := service.NewUserService(log, userRepo)
	postService := service.NewPostService(postRepo, *producer, log)
	likeService := service.NewLikeService(likeRepo, *producer, log)
	followService := service.NewFollowService(followRepo, *producer, log)
	photoService := service.NewPhotoService(s3Repo, log)

	userHandler := handlers.NewUserHandler(userService, log)
	photoHandler := handlers.NewPhotoHandler(userService, photoService, log)
	postHandler := handlers.NewPostHandler(postService, photoService, log)
	LikeHandler := handlers.NewLikeHandler(likeService, log)
	followHandler := handlers.NewFollowHandler(followService, log)

	handler := rest.NewHandler(log, userHandler, photoHandler, postHandler, LikeHandler, followHandler)

	router := handler.InitRouter()

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8082/swagger/doc.json"), // Путь к JSON-файлу Swagger
	))
	srv := &http.Server{
		Addr:    cfg.HttpServe.Address,
		Handler: router,
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
