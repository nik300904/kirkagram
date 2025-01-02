package service

import (
	"log/slog"
)

type PhotoService interface {
	GetPhoto(key string) ([]byte, error)
	UploadPhoto(key string, data []byte) error
}

type Photo struct {
	client PhotoService
	log    *slog.Logger
}

func NewPhotoService(client PhotoService, log *slog.Logger) *Photo {
	return &Photo{
		client: client,
		log:    log,
	}
}

func (p *Photo) UploadPhoto(key string, data []byte) error {
	return p.client.UploadPhoto(key, data)
}

func (p *Photo) GetPhoto(key string) ([]byte, error) {
	return p.client.GetPhoto(key)
}
