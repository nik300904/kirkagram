package service

import (
	"encoding/json"
	"fmt"
	k "kirkagram/internal/kafka"
	"kirkagram/internal/models"
	"log/slog"
)

type LikeService interface {
	LikePostByID(likeReq *models.LikeRequest) error
	UnlikePostByID(likeReq *models.LikeRequest) error
	GetLikesByID(postID int) (models.LikeResponse, error)
}

type Like struct {
	client   LikeService
	producer k.Producer
	log      *slog.Logger
}

func NewLikeService(client LikeService, producer k.Producer, log *slog.Logger) *Like {
	return &Like{
		client:   client,
		producer: producer,
		log:      log,
	}
}

func (l *Like) UnlikePostByID(likeReq *models.LikeRequest) error {
	return l.client.UnlikePostByID(likeReq)
}

func (l *Like) GetLikesByID(postID int) (models.LikeResponse, error) {
	return l.client.GetLikesByID(postID)
}

func (l *Like) LikePostByID(likeReq *models.LikeRequest) error {
	const op = "service.LikePostByID"
	topic := "like"

	likeReqSlc, err := json.Marshal(likeReq)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = l.producer.Produce(likeReqSlc, &topic)
	if err != nil {
		return err
	}

	return l.client.LikePostByID(likeReq)
}
