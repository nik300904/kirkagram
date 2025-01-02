package service

import (
	"kirkagram/internal/models"
	"log/slog"
)

type PostService interface {
	CreatePost(post models.CreatePostRequest) error
}

type Post struct {
	storage PostService
	log     *slog.Logger
}

func NewPostService(storage PostService, log *slog.Logger) PostService {
	return &Post{
		storage: storage,
		log:     log,
	}
}

func (p *Post) CreatePost(post models.CreatePostRequest) error {
	return p.storage.CreatePost(post)
}
