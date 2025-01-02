package rest

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"kirkagram/internal/lib/logger/handlers/customErrors"
	"kirkagram/internal/models"
	"kirkagram/internal/storage"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type User interface {
	GetByEmail(ctx context.Context, email string) (*models.GetUserResponse, error)
	Update(ctx context.Context, updateUser models.UpdateUserRequest) error
	GetAllFollowers(ctx context.Context, userID int) (*[]models.GetAllFollowersResponse, error)
}

type Photo interface {
	GetPhoto(key string) ([]byte, error)
	UploadPhoto(key string, data []byte) error
}

type Handler struct {
	userService  User
	photoService Photo
	log          *slog.Logger
}

func NewHandler(log *slog.Logger, userService User, photoService Photo) *Handler {
	return &Handler{userService: userService, photoService: photoService, log: log}
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
		r.Put("/user", h.UpdateUser)
		r.Get("/user/{userID}/followers", h.GetAllFollowers)
		r.Get("/user/{userID}/following", h.GetAllFollowing)
		r.Get("/photo/{key}", h.GetPhotoURL)
		r.Post("/photo", h.UploadPhoto)
	})

	return router
}

func (h *Handler) UploadPhoto(w http.ResponseWriter, r *http.Request) {
	const op = "transport.rest.UploadPhoto"

	h.log.With(slog.String("op", op))
	h.log.Info("Get photo URL")

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		h.log.Error("Failed to parse multipart form", slog.String("error", err.Error()))

		render.Status(r, http.StatusLengthRequired)
		render.JSON(w, r, customErrors.NewError(err.Error()))

		return
	}

	file, header, err := r.FormFile("photo")
	if err != nil {
		h.log.Error("Failed to get file from form", slog.String("error", err.Error()))

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customErrors.NewError(err.Error()))

		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		h.log.Error("Failed to read file", slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customErrors.NewError(err.Error()))

		return
	}
	filename := header.Filename
	hash := sha256.Sum256([]byte(filename))
	filename = fmt.Sprintf("%x", hash[:8])

	err = h.photoService.UploadPhoto(filename, fileBytes)
	if err != nil {
		h.log.Error("Failed to upload file", slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customErrors.NewError(err.Error()))

		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{"filename": filename})
}

func (h *Handler) GetPhotoURL(w http.ResponseWriter, r *http.Request) {
	const op = "transport.rest.GetPhotoURL"

	h.log.With(slog.String("op", op))
	h.log.Info("Get photo URL")

	key := chi.URLParam(r, "key")

	photo, err := h.photoService.GetPhoto(key)
	if err != nil {
		h.log.Error("Failed to get photo from storage", slog.String("error", err.Error()))

		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, customErrors.NewError(err.Error()))

		return
	}

	h.log.Info("Get photo URL completed successfully")

	render.Status(r, http.StatusOK)
	w.Write(photo)
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

func (h *Handler) GetAllFollowers(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) GetAllFollowing(w http.ResponseWriter, r *http.Request) {
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
