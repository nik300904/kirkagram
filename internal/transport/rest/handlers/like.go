package handlers

import (
	"github.com/go-chi/render"
	"kirkagram/internal/lib/logger/handlers/customResponse"
	"kirkagram/internal/models"
	"log/slog"
	"net/http"
)

type Like interface {
	LikePostByID(likeReq *models.LikeRequest) error
	UnlikePostByID(likeReq *models.LikeRequest) error
}

type LikeHandler struct {
	likeService Like
	log         *slog.Logger
}

func NewLikeHandler(likeService Like, log *slog.Logger) *LikeHandler {
	return &LikeHandler{
		likeService: likeService,
		log:         log,
	}
}

func (l *LikeHandler) UnlikePost(w http.ResponseWriter, r *http.Request) {
	const op = "rest.handlers.like.LikePost"

	log := l.log.With(slog.String("op", op))
	log.Info("starting delete post")

	var req models.LikeRequest
	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		log.Error("unable to decode body", slog.String("error", err.Error()))

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}

	err = l.likeService.UnlikePostByID(&req)
	if err != nil {
		log.Error("unable to like post", slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, customResponse.NewStatus(200))
}

func (l *LikeHandler) LikePost(w http.ResponseWriter, r *http.Request) {
	const op = "rest.handlers.like.LikePost"

	log := l.log.With(slog.String("op", op))
	log.Info("starting delete post")

	var req models.LikeRequest
	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		log.Error("unable to decode body", slog.String("error", err.Error()))

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}

	err = l.likeService.LikePostByID(&req)
	if err != nil {
		log.Error("unable to like post", slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, customResponse.NewStatus(201))
}
