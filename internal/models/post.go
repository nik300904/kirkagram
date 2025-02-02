package models

import "time"

type Posts struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ImageURL  string    `json:"image_url"`
	Caption   string    `json:"caption"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreatePostRequest struct {
	UserID   int    `json:"user_id"`
	Caption  string `json:"caption"`
	ImageURL string `json:"image_url"`
}
