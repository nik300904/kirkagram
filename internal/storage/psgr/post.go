package psgr

import (
	"database/sql"
	"fmt"
	"kirkagram/internal/models"
	"kirkagram/internal/storage"
)

type PostStorage struct {
	db *sql.DB
}

func NewPostStorage(db *sql.DB) *PostStorage {
	return &PostStorage{db: db}
}

func (p *PostStorage) CreatePost(post models.CreatePostRequest) error {
	const op = "storage.psgr.post.CreatePost"

	exec, err := p.db.Exec(`
		INSERT INTO "post" (user_id, image_url, caption)
		VALUES ($1, $2, $3)`, post.UserID, post.ImageURL, post.Caption)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	num, err := exec.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if num == 0 {
		return storage.ErrPostExists
	}

	return nil
}

//func (p *PostStorage) GetAllPosts()
