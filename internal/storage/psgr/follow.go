package psgr

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"kirkagram/internal/models"
	"kirkagram/internal/storage"
)

type FollowStorage struct {
	db *sql.DB
}

func NewFollowStorage(db *sql.DB) *FollowStorage {
	return &FollowStorage{db: db}
}

func (f *FollowStorage) FollowByID(req models.FollowRequest) error {
	const op = "storage.psgr.follow.FollowByID"

	if req.FollowerID == req.FollowingID {
		return fmt.Errorf("%s: %w", op, storage.SelfFollowError)
	}

	_, err := f.db.Exec(`
		INSERT INTO "follow" ("follower_id", "following_id")
		VALUES ($1, $2)
		`,
		req.FollowerID,
		req.FollowingID,
	)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Message {
			case "insert or update on table \"follow\" violates foreign key constraint \"follow_follower_id_fkey\"":
				return fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
			case "duplicate key value violates unique constraint \"follow_follower_id_following_id_key\"":
				return fmt.Errorf("%s: %w", op, storage.ErrAlreadyFollowed)
			}
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (f *FollowStorage) UnFollowByID(req models.FollowRequest) error {
	const op = "storage.psgr.follow.UnFollow"

	if req.FollowerID == req.FollowingID {
		return fmt.Errorf("%s: %w", op, storage.SelfUnFollowError)
	}

	_, err := f.db.Exec(
		`DELETE FROM "follow" WHERE follower_id = $1 AND following_id = $2`,
		req.FollowerID,
		req.FollowingID,
	)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Message {
			case "insert or update on table \"follow\" violates foreign key constraint \"follow_follower_id_fkey\"":
				return fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
			case "duplicate key value violates unique constraint \"follow_follower_id_following_id_key\"":
				return fmt.Errorf("%s: %w", op, storage.ErrAlreadyFollowed)
			}
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
