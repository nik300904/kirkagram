package models

import "time"

type Like struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
}

type LikeRequest struct {
	UserID int `json:"user_id"`
	PostID int `json:"post_id"`
}
