package rest

import (
	"context"
	"errors"
	"kirkagram/internal/lib/logger/handlers/customErrors"
	"kirkagram/internal/models"
	"kirkagram/internal/storage"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type User interface {
	GetByEmail(ctx context.Context, email string) (*models.GetUserResponse, error)
	Update(ctx context.Context, updateUser models.UpdateUserRequest) error
}

type Handler struct {
	userService User
	log         *slog.Logger
}

func NewHandler(log *slog.Logger, userService User) *Handler {
	return &Handler{userService: userService, log: log}
}

func (h *Handler) InitRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		h.log.Info("Hello")
		w.Write([]byte("Hello"))
	})

	router.Route("/api", func(r chi.Router) {
		h.log.Info("Init api routes")
		r.Get("/user/{email}", h.GetUser)
	})

	return router
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Get user")

	ctx := context.Background()
	email := chi.URLParam(r, "email")

	h.log.Info("Get user by email", slog.String("email", email))

	user, err := h.userService.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			h.log.Error("Get user by email with error", slog.String("email", email), slog.String("error", err.Error()))

			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, customErrors.NewError(err.Error()))

			return
		}

		h.log.Error("Get user by email with error", slog.String("email", email), slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customErrors.NewError(err.Error()))

		return
	}

	h.log.Info("Get user by email completed", slog.String("email", email))

	render.JSON(w, r, user)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Update user")

	ctx := context.Background()

	var updateUser models.UpdateUserRequest
	if err := render.DecodeJSON(r.Body, &updateUser); err != nil {
		h.log.Error("Update user with error", slog.String("error", err.Error()))

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customErrors.NewError(err.Error()))

		return
	}

	err := h.userService.Update(ctx, updateUser)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			h.log.Error("Update user with error", slog.String("error", err.Error()))

			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, customErrors.NewError(err.Error()))

			return
		}

		h.log.Error("Update user with error", slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customErrors.NewError(err.Error()))

		return
	}

	h.log.Info("Update user completed")

	render.Status(r, http.StatusOK)
	render.JSON(w, r, nil)
}
