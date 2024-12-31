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

func (s *UserStorage) GetAllFollowers(userID int) (*[]models.GetAllFollowersResponse, error) {
	const op = "storage.postgres.GetAllFollowers"

	rows, err := s.db.Query(`
		SELECT DISTINCT u.username, u.profile_pic
		FROM "user" main_user
		CROSS JOIN LATERAL unnest(main_user.followers) AS follower_id
		JOIN "user" u ON u.id = follower_id
		WHERE main_user.id = $1;
	`, userID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrUserNotFound
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var followers []models.GetAllFollowersResponse

	for rows.Next() {
		var follower models.GetAllFollowersResponse

		if err := rows.Scan(&follower.Username, &follower.ProfilePic); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		followers = append(followers, follower)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &followers, nil
}

func (s *UserStorage) GetAllFollowing(userID int) (*[]models.GetAllFollowersResponse, error) {
	const op = "storage.postgres.GetAllFollowing"

	rows, err := s.db.Query(`
		SELECT DISTINCT u.username, u.profile_pic
		FROM "user" main_user
		CROSS JOIN LATERAL unnest(main_user.followers) AS follower_id
		JOIN "user" u ON u.id = follower_id
		WHERE main_user.id = $1;
	`, userID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrUserNotFound
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var followers []models.GetAllFollowersResponse

	for rows.Next() {
		var follower models.GetAllFollowersResponse

		if err := rows.Scan(&follower.Username, &follower.ProfilePic); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		followers = append(followers, follower)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &followers, nil
}
