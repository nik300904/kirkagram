package handlers

import (
	"crypto/sha256"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"io"
	"kirkagram/internal/lib/logger/handlers/customResponse"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

type Photo interface {
	GetPhoto(key string) ([]byte, error)
	UploadPhoto(key string, data []byte) error
}

type UserForPhoto interface {
	UploadProfilePic(userID int, filename string) error
}

type PhotoHandler struct {
	userService  UserForPhoto
	photoService Photo
	log          *slog.Logger
}

func NewPhotoHandler(userService UserForPhoto, photoService Photo, log *slog.Logger) *PhotoHandler {
	return &PhotoHandler{
		userService:  userService,
		photoService: photoService,
		log:          log,
	}
}

func (h *PhotoHandler) UploadPhoto(w http.ResponseWriter, r *http.Request) {
	const op = "rest.handlers.photo.UploadPhoto"

	log := h.log.With(slog.String("op", op))
	log.Info("Get photo URL")

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Error("Failed to parse multipart form", slog.String("error", err.Error()))

		render.Status(r, http.StatusLengthRequired)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}

	file, header, err := r.FormFile("photo")
	if err != nil {
		log.Error("Failed to get file from form", slog.String("error", err.Error()))

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Error("Failed to read file", slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}
	filename := header.Filename
	currentTime := time.Now()
	timestampString := currentTime.Format("2006-01-02_15-04-05.000000")
	filename = filename + timestampString
	hash := sha256.Sum256([]byte(filename))
	filename = fmt.Sprintf("%x", hash[:8])

	userID := r.FormValue("id")
	num, err := strconv.Atoi(userID)
	if err != nil {
		log.Error("Failed to convert user ID to int", slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}

	err = h.userService.UploadProfilePic(num, filename)
	if err != nil {
		log.Error("Failed to upload file to bd", slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}

	err = h.photoService.UploadPhoto(filename, fileBytes)
	if err != nil {
		log.Error("Failed to upload file", slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{"filename": filename})
}

func (h *PhotoHandler) GetPhotoURL(w http.ResponseWriter, r *http.Request) {
	const op = "rest.handlers.photo.GetPhotoURL"

	log := h.log.With(slog.String("op", op))
	log.Info("Get photo URL")

	key := chi.URLParam(r, "key")

	photo, err := h.photoService.GetPhoto(key)
	if err != nil {
		log.Error("Failed to get photo from storage", slog.String("error", err.Error()))

		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}

	log.Info("Get photo URL completed successfully")

	render.Status(r, http.StatusOK)
	w.Write(photo)
}
