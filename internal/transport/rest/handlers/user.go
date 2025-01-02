package handlers

import (
	"context"
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"kirkagram/internal/lib/logger/handlers/customErrors"
	"kirkagram/internal/models"
	"kirkagram/internal/storage"
	"log/slog"
	"net/http"
	"strconv"
)

type User interface {
	GetByID(ctx context.Context, ID string) (*models.GetUserResponse, error)
	Update(ctx context.Context, updateUser models.UpdateUserRequest) error
	GetAllFollowers(ctx context.Context, userID int) (*[]models.GetAllFollowersResponse, error)
	UploadProfilePic(userID int, filename string) error
}

type UserHandler struct {
	userService User
	log         *slog.Logger
}

func NewUserHandler(userService User, log *slog.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		log:         log,
	}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Get user")

	ctx := context.Background()
	id := chi.URLParam(r, "id")

	h.log.Info("Get user by email", slog.String("email", id))

	user, err := h.userService.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			h.log.Error("Get user by email with error", slog.String("email", id), slog.String("error", err.Error()))

			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, customErrors.NewError(err.Error()))

			return
		}

		h.log.Error("Get user by email with error", slog.String("email", id), slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customErrors.NewError(err.Error()))

		return
	}

	h.log.Info("Get user by email completed", slog.String("email", id))

	render.JSON(w, r, user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
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

func (h *UserHandler) GetAllFollowers(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Get all followers")

	ctx := context.Background()
	userID := chi.URLParam(r, "userID")

	if userID == "" {
		h.log.Error("Get all followers with error", slog.String("error", "userID is empty"))

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customErrors.NewError("userID is empty"))

		return
	}

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		h.log.Error("Get all followers with error", slog.String("userID", userID), slog.String("error", err.Error()))

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customErrors.NewError(err.Error()))

		return
	}

	followers, err := h.userService.GetAllFollowers(ctx, userIDInt)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			h.log.Error("Get all followers with error", slog.Int("userID", userIDInt), slog.String("error", err.Error()))

			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, customErrors.NewError(err.Error()))

			return
		}

		h.log.Error("Get all followers with error", slog.Int("userID", userIDInt), slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customErrors.NewError(err.Error()))

		return
	}

	h.log.Info("Get all followers completed", slog.Int("userID", userIDInt))

	render.Status(r, http.StatusOK)
	render.JSON(w, r, followers)
}

func (h *UserHandler) GetAllFollowing(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Get all following")

	ctx := context.Background()
	userID := chi.URLParam(r, "userID")

	if userID == "" {
		h.log.Error("Get all following with error", slog.String("error", "userID is empty"))

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customErrors.NewError("userID is empty"))

		return
	}

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		h.log.Error("Get all following with error", slog.String("userID", userID), slog.String("error", err.Error()))

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customErrors.NewError(err.Error()))

		return
	}

	followers, err := h.userService.GetAllFollowers(ctx, userIDInt)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			h.log.Error("Get all following with error", slog.Int("userID", userIDInt), slog.String("error", err.Error()))

			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, customErrors.NewError(err.Error()))

			return
		}

		h.log.Error("Get all following with error", slog.Int("userID", userIDInt), slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customErrors.NewError(err.Error()))

		return
	}

	h.log.Info("Get all following completed", slog.Int("userID", userIDInt))

	render.Status(r, http.StatusOK)
	render.JSON(w, r, followers)
}
