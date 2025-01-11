package handlers

import (
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"kirkagram/internal/lib/logger/handlers/customResponse"
	"kirkagram/internal/models"
	"log/slog"
	"net/http"
	"strconv"
)

type Like interface {
	LikePostByID(likeReq *models.LikeRequest) error
	UnlikePostByID(likeReq *models.LikeRequest) error
	GetLikesByID(postID int) (models.LikeResponse, error)
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

// GetLikes godoc
// @Summary Get likes count for a post
// @Description Get the number of likes for a specific post
// @Tags likes
// @Accept json
// @Produce json
// @Param postID path int true "Post ID"
// @Success 200 {object} models.LikeResponse
// @Failure 400 {object} customResponse.Error
// @Failure 500 {object} customResponse.Error
// @Router /like/{postID} [get]
func (l *LikeHandler) GetLikes(w http.ResponseWriter, r *http.Request) {
	const op = "rest.handlers.like.LikePost"

	log := l.log.With(slog.String("op", op))
	log.Info("starting delete post")

	postID := chi.URLParam(r, "postID")
	if postID == "" {
		log.Error("postID is required")

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customResponse.NewError("postID empty"))

		return
	}

	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		log.Error("invalid postID", slog.String("postID", postID))

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customResponse.NewError("postID must be numeric"))

		return
	}

	count, err := l.likeService.GetLikesByID(postIDInt)
	if err != nil {
		log.Error("error getting likes by postID", slog.String("postID", postID), slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		originalErr := errors.Unwrap(err)
		render.JSON(w, r, customResponse.NewError(originalErr.Error()))

		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, count)
}

// UnlikePost godoc
// @Summary Unlike a post
// @Description Remove a like from a specific post
// @Tags likes
// @Accept json
// @Produce json
// @Param request body models.LikeRequest true "Unlike request"
// @Success 200 {object} customResponse.CustomStatus
// @Failure 400 {object} customResponse.Error
// @Failure 500 {object} customResponse.Error
// @Router /like [delete]
func (l *LikeHandler) UnlikePost(w http.ResponseWriter, r *http.Request) {
	const op = "rest.handlers.like.LikePost"

	log := l.log.With(slog.String("op", op))
	log.Info("starting delete post")

	var req models.LikeRequest
	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		log.Error("unable to decode body", slog.String("error", err.Error()))

		render.Status(r, http.StatusBadRequest)
		originalErr := errors.Unwrap(err)
		render.JSON(w, r, customResponse.NewError(originalErr.Error()))

		return
	}

	err = l.likeService.UnlikePostByID(&req)
	if err != nil {
		log.Error("unable to like post", slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		originalErr := errors.Unwrap(err)
		render.JSON(w, r, customResponse.NewError(originalErr.Error()))

		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, customResponse.NewStatus(200))
}

// LikePost godoc
// @Summary Like a post
// @Description Add a like to a specific post
// @Tags likes
// @Accept json
// @Produce json
// @Param request body models.LikeRequest true "Like request"
// @Success 201 {object} customResponse.CustomStatus
// @Failure 400 {object} customResponse.Error
// @Failure 500 {object} customResponse.Error
// @Router /like [post]
func (l *LikeHandler) LikePost(w http.ResponseWriter, r *http.Request) {
	const op = "rest.handlers.like.LikePost"

	log := l.log.With(slog.String("op", op))
	log.Info("starting delete post")

	var req models.LikeRequest
	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		log.Error("unable to decode body", slog.String("error", err.Error()))

		render.Status(r, http.StatusBadRequest)
		originalErr := errors.Unwrap(err)
		render.JSON(w, r, customResponse.NewError(originalErr.Error()))

		return
	}

	err = l.likeService.LikePostByID(&req)
	if err != nil {
		log.Error("unable to like post", slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		originalErr := errors.Unwrap(err)
		render.JSON(w, r, customResponse.NewError(originalErr.Error()))

		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, customResponse.NewStatus(201))
}
