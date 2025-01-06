package service

import (
	"kirkagram/internal/models"
	"log/slog"
)

type LikeService interface {
	LikePostByID(likeReq *models.LikeRequest) error
}

type Like struct {
	client LikeService
	log    *slog.Logger
}

func NewLikeService(client LikeService, log *slog.Logger) *Like {
	return &Like{
		client: client,
		log:    log,
	}
}

func (l *Like) LikePostByID(likeReq *models.LikeRequest) error {
	return l.client.LikePostByID(likeReq)
}
