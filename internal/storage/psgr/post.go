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

func (p *PostStorage) GetAllPosts() (*[]models.Posts, error) {
	const op = "storage.psgr.post.GetAllPosts"

	rows, err := p.db.Query("SELECT * FROM post")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var posts []models.Posts

	for rows.Next() {
		var post models.Posts

		if err := rows.Scan(&post.ID, &post.UserID, &post.ImageURL, &post.Caption, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &posts, nil
}
