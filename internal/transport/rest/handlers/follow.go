package handlers

import (
	"errors"
	"github.com/go-chi/render"
	"kirkagram/internal/lib/logger/handlers/customResponse"
	"kirkagram/internal/models"
	"log/slog"
	"net/http"
)

type Follow interface {
	FollowByID(req models.FollowRequest) error
	UnFollowByID(req models.FollowRequest) error
}

type FollowHandler struct {
	followService Follow
	log           *slog.Logger
}

func NewFollowHandler(followService Follow, log *slog.Logger) *FollowHandler {
	return &FollowHandler{
		followService: followService,
		log:           log,
	}
}

func (f *FollowHandler) Follow(w http.ResponseWriter, r *http.Request) {
	const op = "rest.handlers.follow.Follow"

	log := f.log.With(slog.String("op", op))
	log.Info("starting delete post")

	var req models.FollowRequest

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		log.Error("error decoding body", slog.String("error", err.Error()))

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}

	err = f.followService.FollowByID(req)
	if err != nil {
		log.Error("error following post", slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		originalErr := errors.Unwrap(err)
		render.JSON(w, r, customResponse.NewError(originalErr.Error()))

		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, customResponse.NewStatus(201))
}

func (f *FollowHandler) UnFollow(w http.ResponseWriter, r *http.Request) {
	const op = "rest.handlers.follow.UnFollow"

	log := f.log.With(slog.String("op", op))
	log.Info("starting delete post")

	var req models.FollowRequest

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		log.Error("error decoding body", slog.String("error", err.Error()))

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}

	err = f.followService.UnFollowByID(req)
	if err != nil {
		log.Error("error following post", slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		originalErr := errors.Unwrap(err)
		render.JSON(w, r, customResponse.NewError(originalErr.Error()))

		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, customResponse.NewStatus(201))
}
