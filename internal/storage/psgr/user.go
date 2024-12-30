package psgr

import (
	"database/sql"
	"errors"
	"fmt"
	"kirkagram/internal/models"
	"kirkagram/internal/storage"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (s *UserStorage) GetByEmail(email string) (*models.GetUserResponse, error) {
	const op = "storage.postgres.GetUser"

	var user models.GetUserResponse

	row := s.db.QueryRow(`SELECT "email", username FROM "user" WHERE "email" = $1`, email).Scan(&user.Email, &user.Username)

	if row != nil {
		if errors.Is(row, sql.ErrNoRows) {
			return nil, storage.ErrUserNotFound
		}

		return nil, fmt.Errorf("%s: %w", op, row)
	}

	return &user, nil
}
