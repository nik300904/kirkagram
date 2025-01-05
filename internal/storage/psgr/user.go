package psgr

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"kirkagram/internal/models"
	"kirkagram/internal/storage"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (s *UserStorage) CreateUser(user *models.CreateUserRequest) error {
	const op = "storage.psgr.user.CreateUser"

	exec, err := s.db.Exec(
		`INSERT INTO "user" (username, email, password) VALUES ($1, $2, $3)`,
		user.Username,
		user.Email,
		user.Password,
	)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			switch err.Message {
			case "duplicate key value violates unique constraint \"user_email_key\"":
				return storage.ErrEmailAlreadyRegistered
			case "duplicate key value violates unique constraint \"user_username_key\"":
				return storage.ErrUsernameAlreadyRegistered
			}
		}

		return err
	}

	num, err := exec.RowsAffected()
	if err != nil {
		fmt.Println("%s: %w", op, err)
		return err
	}

	if num == 0 {
		return storage.ErrUserAlreadyExists
	}

	return nil
}

func (s *UserStorage) DeleteUser(ID int64) error {
	const op = "storage.psgr.user.DeleteUser"

	exec, err := s.db.Exec(
		`DELETE FROM "user" WHERE id=$1`,
		ID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return storage.ErrUserNotFound
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	num, err := exec.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if num == 0 {
		return storage.ErrUserNotFound
	}

	return nil
}

func (s *UserStorage) UploadProfilePic(userID int, filename string) error {
	const op = "storage.psgr.user.UploadProfilePic"

	profilePic := fmt.Sprintf("api/photo/%v", filename)

	exec, err := s.db.Exec(
		`UPDATE "user" SET "profile_pic" = $1 WHERE "id" = $2`,
		profilePic,
		userID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	count, err := exec.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if count == 0 {
		return fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
	}

	return nil
}

func (s *UserStorage) GetByID(ID string) (*models.GetUserResponse, error) {
	const op = "storage.psgr.user.GetUser"

	var user models.GetUserResponse

	row := s.db.QueryRow(
		`SELECT "id", "email", "username", "bio", "profile_pic" FROM "user" WHERE "id" = $1`,
		ID,
	).Scan(&user.ID, &user.Email, &user.Username, &user.Bio, &user.ProfilePic)

	if row != nil {
		if errors.Is(row, sql.ErrNoRows) {
			return nil, storage.ErrUserNotFound
		}

		return nil, fmt.Errorf("%s: %w", op, row)
	}

	return &user, nil
}

func (s *UserStorage) Update(updateUser models.UpdateUserRequest) error {
	const op = "storage.psgr.user.Update"

	row, err := s.db.Exec(
		`UPDATE "user" SET "username" = $1, "email" = $2, "bio" = $3 WHERE "id" = $5`,
		updateUser.Username,
		updateUser.Email,
		updateUser.Bio,
		updateUser.ID,
	)

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
	const op = "storage.psgr.user.GetAllFollowers"

	rows, err := s.db.Query(
		`
		SELECT DISTINCT u.username, u.profile_pic
		FROM "user" main_user
		CROSS JOIN LATERAL unnest(main_user.followers) AS follower_id
		JOIN "user" u ON u.id = follower_id
		WHERE main_user.id = $1;
	`,
		userID,
	)

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
	const op = "storage.psgr.user.GetAllFollowing"

	rows, err := s.db.Query(
		`
		SELECT DISTINCT u.username, u.profile_pic
		FROM "user" main_user
		CROSS JOIN LATERAL unnest(main_user.followers) AS follower_id
		JOIN "user" u ON u.id = follower_id
		WHERE main_user.id = $1;
		`,
		userID,
	)

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
