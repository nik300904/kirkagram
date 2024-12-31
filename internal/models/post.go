package models

import "time"

type Posts struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	ImageURL  string    `json:"image_url"`
	Caption   string    `json:"caption"`
	Likes     int       `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
}
