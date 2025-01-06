package psgr

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"kirkagram/internal/models"
	"kirkagram/internal/storage"
)

type LikeStorage struct {
	db *sql.DB
}

func NewLikeStorage(db *sql.DB) *LikeStorage {
	return &LikeStorage{db: db}
}

func (l *LikeStorage) UnlikePostByID(likeReq *models.LikeRequest) error {
	const op = "storage.psgr.like.LikePostByID"

	exec, err := l.db.Exec(
		`DELETE FROM "like" WHERE user_id = $1 AND post_id = $2`,
		likeReq.UserID,
		likeReq.PostID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return storage.ErrLikeNotFound
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	num, err := exec.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if num == 0 {
		return fmt.Errorf("%s: %w", op, storage.ErrLikeNotFound)
	}

	return nil
}

func (l *LikeStorage) LikePostByID(likeReq *models.LikeRequest) error {
	const op = "storage.psgr.like.LikePostByID"

	exec, err := l.db.Exec(
		`INSERT INTO "like" (user_id, post_id) VALUES ($1, $2)`,
		likeReq.UserID,
		likeReq.PostID,
	)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Message == "duplicate key value violates unique constraint \"like_user_id_post_id_key\"" {
				return storage.ErrPostAlreadyLiked
			}
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	num, err := exec.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if num == 0 {
		return storage.ErrPostAlreadyLiked
	}

	return nil
}
