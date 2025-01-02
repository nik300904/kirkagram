package service

import (
	"kirkagram/internal/models"
	"log/slog"
)

type PostService interface {
	CreatePost(post models.CreatePostRequest) error
	GetAllPosts() (*[]models.Posts, error)
	GetPostByID(ID int64) (*models.Posts, error)
}

type Post struct {
	storage PostService
	log     *slog.Logger
}

func NewPostService(storage PostService, log *slog.Logger) *Post {
	return &Post{
		storage: storage,
		log:     log,
	}
}

func (p *Post) CreatePost(post models.CreatePostRequest) error {
	return p.storage.CreatePost(post)
}

func (p *Post) GetAllPosts() (*[]models.Posts, error) {
	return p.storage.GetAllPosts()
}

func (p *Post) GetPostByID(ID int64) (*models.Posts, error) {
	return p.storage.GetPostByID(ID)
}
