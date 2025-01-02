package service

import (
	"context"
	"errors"
	"fmt"
	"kirkagram/internal/models"
	"kirkagram/internal/storage"
	"log/slog"

	"github.com/go-playground/validator"
)

type UserService interface {
	GetByID(ID string) (*models.GetUserResponse, error)
	Update(updateUser models.UpdateUserRequest) error
	GetAllFollowers(userID int) (*[]models.GetAllFollowersResponse, error)
	GetAllFollowing(userID int) (*[]models.GetAllFollowersResponse, error)
	UploadProfilePic(userID int, filename string) error
}

type User struct {
	storage UserService
	log     *slog.Logger
}

func NewUserService(log *slog.Logger, storage UserService) *User {
	return &User{storage: storage, log: log}
}

func (s *User) UploadProfilePic(userID int, filename string) error {
	return s.storage.UploadProfilePic(userID, filename)
}

func (s *User) GetByID(ctx context.Context, ID string) (*models.GetUserResponse, error) {
	const op = "service.user.GetByEmail"

	return s.storage.GetByID(ID)
}

func (s *User) Update(ctx context.Context, updateUser models.UpdateUserRequest) error {
	const op = "service.user.Update"

	validate := validator.New()
	emailStr := models.GetUserValidate{Email: updateUser.Email}

	err := validate.Struct(emailStr)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			s.log.Error(fmt.Sprintf("Field: %s, Tag: %s\n", err.Field(), err.Tag()))
		}

		return fmt.Errorf("%s: %w", op, models.ErrEmailValidate)
	}

	err = s.storage.Update(updateUser)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			s.log.Error("Get user by email with error", slog.String("email", updateUser.Email), slog.String("error", err.Error()))

			return fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *User) GetAllFollowers(ctx context.Context, userID int) (*[]models.GetAllFollowersResponse, error) {
	const op = "service.user.GetAllFollowers"

	followers, err := s.storage.GetAllFollowers(userID)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			s.log.Error("Get followers by userID with error", slog.Int("userID", userID), slog.String("error", err.Error()))

			return nil, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}
	}

	return followers, nil
}

func (s *User) GetAllFollowing(ctx context.Context, userID int) (*[]models.GetAllFollowersResponse, error) {
	const op = "service.user.GetAllFollowing"

	following, err := s.storage.GetAllFollowing(userID)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			s.log.Error("Get followers by userID with error", slog.Int("userID", userID), slog.String("error", err.Error()))

			return nil, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}
	}

	return following, nil
}
