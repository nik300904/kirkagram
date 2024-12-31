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

func (s *UserStorage) Update(updateUser models.UpdateUserRequest) error {
	const op = "storage.postgres.Update"

	row, err := s.db.Exec(`UPDATE "user" SET "username" = $1, "email" = $2, "bio" = $3 WHERE "id" = $5`, updateUser.Username, updateUser.Email, updateUser.Bio, updateUser.ID)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	affectedRows, err := row.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if affectedRows == 0 {
		return fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
	}

	return nil
}
