package service

import (
	"encoding/json"
	"fmt"
	k "kirkagram/internal/kafka"
	"kirkagram/internal/models"
	"log/slog"
)

type PostService interface {
	CreatePost(post models.CreatePostRequest) error
	GetAllPosts() (*[]models.Posts, error)
	GetPostByID(ID int64) (*models.Posts, error)
	GetAllPostsByUserID(userID int64) (*[]models.Posts, error)
	DeletePost(ID int64) error
}

type Post struct {
	storage  PostService
	producer k.Producer
	log      *slog.Logger
}

func NewPostService(storage PostService, producer k.Producer, log *slog.Logger) *Post {
	return &Post{
		storage:  storage,
		producer: producer,
		log:      log,
	}
}

func (p *Post) DeletePost(ID int64) error {
	return p.storage.DeletePost(ID)
}

func (p *Post) GetAllPostsByUserID(userID int64) (*[]models.Posts, error) {
	return p.storage.GetAllPostsByUserID(userID)
}

func (p *Post) CreatePost(post models.CreatePostRequest) error {
	const op = "service.CreatePost"
	topic := "post"

	postSlc, err := json.Marshal(post)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = p.producer.Produce(postSlc, &topic)
	if err != nil {
		return err
	}

	return p.storage.CreatePost(post)
}

func (p *Post) GetAllPosts() (*[]models.Posts, error) {
	return p.storage.GetAllPosts()
}

func (p *Post) GetPostByID(ID int64) (*models.Posts, error) {
	return p.storage.GetPostByID(ID)
}
