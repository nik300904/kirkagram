package handlers

import (
	"context"
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"kirkagram/internal/lib/logger/handlers/customResponse"
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
	GetAllFollowing(ctx context.Context, userID int) (*[]models.GetAllFollowersResponse, error)
	UploadProfilePic(userID int, filename string) error
	DeleteUser(ID int64) error
	RegisterUser(user models.CreateUserRequest) error
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

// Register godoc
// @Summary Register a new user
// @Description Register a new user in the system
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User registration details"
// @Success 201 {object} customResponse.CustomStatus
// @Failure 400 {object} customResponse.Error
// @Failure 500 {object} customResponse.Error
// @Router /user [post]
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	const op = "rest.handlers.user.Register"

	log := h.log.With(slog.String("op", op))
	log.Info("start register user")

	var user models.CreateUserRequest
	err := render.DecodeJSON(r.Body, &user)
	if err != nil {
		log.Error("decode json error", slog.String("error", err.Error()))

		render.Status(r, http.StatusBadRequest)
		originalErr := errors.Unwrap(err)
		render.JSON(w, r, customResponse.NewError(originalErr.Error()))

		return
	}

	err = h.userService.RegisterUser(user)
	if err != nil {
		log.Error("register user error", slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		originalErr := errors.Unwrap(err)
		render.JSON(w, r, customResponse.NewError(originalErr.Error()))

		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, customResponse.NewStatus(201))
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user from the system by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} customResponse.CustomStatus
// @Failure 400 {object} customResponse.Error
// @Failure 404 {object} customResponse.Error
// @Failure 500 {object} customResponse.Error
// @Router /user/{id} [delete]
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	const op = "rest.handlers.user.DeleteUser"

	ID := chi.URLParam(r, "Id")

	log := h.log.With(slog.String("op", op), slog.String("ID", ID))
	log.Info("start delete user")

	if ID == "" {
		log.Error("id is empty")

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customResponse.NewError("id is empty"))

		return
	}

	num, err := strconv.Atoi(ID)
	if err != nil {
		log.Error("id is invalid")

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customResponse.NewError("id must be numeric"))

		return
	}

	err = h.userService.DeleteUser(int64(num))
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			log.Error("user not found")

			render.Status(r, http.StatusNotFound)
			originalErr := errors.Unwrap(err)
			render.JSON(w, r, customResponse.NewError(originalErr.Error()))

			return
		}

		log.Error(err.Error())

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, customResponse.NewStatus(200))
}

// GetUser godoc
// @Summary Get user by ID
// @Description Get details of a specific user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} customResponse.Error
// @Failure 404 {object} customResponse.Error
// @Router /user/{id} [get]
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	const op = "rest.handlers.user.GetUser"

	log := h.log.With(slog.String("op", op))
	log.Info("Get user")

	ctx := context.Background()
	ID := chi.URLParam(r, "id")

	if ID == "" {
		log.Error("id is empty")

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customResponse.NewError("id is empty"))

		return
	}

	log.Info("Get user by email", slog.String("id", ID))

	user, err := h.userService.GetByID(ctx, ID)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			log.Error("Get user by email with error", slog.String("id", ID), slog.String("error", err.Error()))

			render.Status(r, http.StatusNotFound)
			originalErr := errors.Unwrap(err)
			render.JSON(w, r, customResponse.NewError(originalErr.Error()))

			return
		}

		log.Error("Get user by email with error", slog.String("id", ID), slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		originalErr := errors.Unwrap(err)
		render.JSON(w, r, customResponse.NewError(originalErr.Error()))

		return
	}

	log.Info("Get user by email completed", slog.String("id", ID))

	render.Status(r, http.StatusOK)
	render.JSON(w, r, user)
}

// UpdateUser godoc
// @Summary Update user information
// @Description Update details of an existing user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UpdateUserRequest true "Updated user information"
// @Success 200 {object} customResponse.CustomStatus
// @Failure 400 {object} customResponse.Error
// @Failure 404 {object} customResponse.Error
// @Failure 500 {object} customResponse.Error
// @Router /user [put]
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	const op = "rest.handlers.user.UpdateUser"

	log := h.log.With(slog.String("op", op))
	log.Info("Update user")

	ctx := context.Background()

	var updateUser models.UpdateUserRequest
	if err := render.DecodeJSON(r.Body, &updateUser); err != nil {
		log.Error("Update user with error", slog.String("error", err.Error()))

		render.Status(r, http.StatusBadRequest)
		originalErr := errors.Unwrap(err)
		render.JSON(w, r, customResponse.NewError(originalErr.Error()))

		return
	}

	err := h.userService.Update(ctx, updateUser)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			log.Error("Update user with error", slog.String("error", err.Error()))

			render.Status(r, http.StatusNotFound)
			originalErr := errors.Unwrap(err)
			render.JSON(w, r, customResponse.NewError(originalErr.Error()))

			return
		}

		log.Error("Update user with error", slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		originalErr := errors.Unwrap(err)
		render.JSON(w, r, customResponse.NewError(originalErr.Error()))

		return
	}

	log.Info("Update user completed")

	render.Status(r, http.StatusOK)
	render.JSON(w, r, customResponse.NewStatus(200))
}

// GetAllFollowers godoc
// @Summary Get all followers of a user
// @Description Retrieve a list of all followers for a specific user
// @Tags users
// @Accept json
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {array} models.GetAllFollowersResponse
// @Failure 400 {object} customResponse.Error
// @Failure 404 {object} customResponse.Error
// @Failure 500 {object} customResponse.Error
// @Router /user/{userID}/followers [get]
func (h *UserHandler) GetAllFollowers(w http.ResponseWriter, r *http.Request) {
	const op = "rest.handlers.user.GetAllFollowers"

	log := h.log.With(slog.String("op", op))
	log.Info("Get all followers")

	ctx := context.Background()
	userID := chi.URLParam(r, "userID")

	if userID == "" {
		log.Error("Get all followers with error", slog.String("error", "userID is empty"))

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customResponse.NewError("userID is empty"))

		return
	}

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		log.Error("Get all followers with error", slog.String("userID", userID), slog.String("error", err.Error()))

		render.Status(r, http.StatusBadRequest)
		originalErr := errors.Unwrap(err)
		render.JSON(w, r, customResponse.NewError(originalErr.Error()))

		return
	}

	followers, err := h.userService.GetAllFollowers(ctx, userIDInt)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			log.Error("Get all followers with error", slog.Int("userID", userIDInt), slog.String("error", err.Error()))

			render.Status(r, http.StatusNotFound)
			originalErr := errors.Unwrap(err)
			render.JSON(w, r, customResponse.NewError(originalErr.Error()))

			return
		}

		log.Error("Get all followers with error", slog.Int("userID", userIDInt), slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		originalErr := errors.Unwrap(err)
		render.JSON(w, r, customResponse.NewError(originalErr.Error()))

		return
	}

	log.Info("Get all followers completed", slog.Int("userID", userIDInt))

	render.Status(r, http.StatusOK)
	render.JSON(w, r, followers)
}

// GetAllFollowing godoc
// @Summary Get all users followed by a user
// @Description Retrieve a list of all users followed by a specific user
// @Tags users
// @Accept json
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {array} models.GetAllFollowersResponse
// @Failure 400 {object} customResponse.Error
// @Failure 404 {object} customResponse.Error
// @Failure 500 {object} customResponse.Error
// @Router /user/{userID}/following [get]
func (h *UserHandler) GetAllFollowing(w http.ResponseWriter, r *http.Request) {
	const op = "rest.handlers.user.GetAllFollowing"

	log := h.log.With(slog.String("op", op))
	log.Info("Get all following")

	ctx := context.Background()
	userID := chi.URLParam(r, "userID")

	if userID == "" {
		log.Error("Get all following with error", slog.String("error", "userID is empty"))

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customResponse.NewError("userID is empty"))

		return
	}

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		log.Error("Get all following with error", slog.String("userID", userID), slog.String("error", err.Error()))

		render.Status(r, http.StatusBadRequest)
		originalErr := errors.Unwrap(err)
		render.JSON(w, r, customResponse.NewError(originalErr.Error()))

		return
	}

	followers, err := h.userService.GetAllFollowing(ctx, userIDInt)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			log.Error("Get all following with error", slog.Int("userID", userIDInt), slog.String("error", err.Error()))

			render.Status(r, http.StatusNotFound)
			originalErr := errors.Unwrap(err)
			render.JSON(w, r, customResponse.NewError(originalErr.Error()))

			return
		}

		log.Error("Get all following with error", slog.Int("userID", userIDInt), slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		originalErr := errors.Unwrap(err)
		render.JSON(w, r, customResponse.NewError(originalErr.Error()))

		return
	}

	log.Info("Get all following completed", slog.Int("userID", userIDInt))

	render.Status(r, http.StatusOK)
	render.JSON(w, r, followers)
}
