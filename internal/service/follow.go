package service

import (
	"encoding/json"
	"fmt"
	k "kirkagram/internal/kafka"
	"kirkagram/internal/models"
	"log/slog"
)

type FollowService interface {
	FollowByID(req models.FollowRequest) error
	UnFollowByID(req models.FollowRequest) error
}

type Follow struct {
	client   FollowService
	producer k.Producer
	log      *slog.Logger
}

func NewFollowService(client FollowService, producer k.Producer, log *slog.Logger) *Follow {
	return &Follow{
		client:   client,
		producer: producer,
		log:      log,
	}
}

func (f *Follow) FollowByID(req models.FollowRequest) error {
	const op = "service.follow.FollowByID"

	if err := f.client.FollowByID(req); err != nil {
		return err
	}

	followReqSlc, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = f.producer.Produce(followReqSlc, "follow")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (f *Follow) UnFollowByID(req models.FollowRequest) error {
	const op = "service.follow.UnFollowByID"

	if err := f.client.UnFollowByID(req); err != nil {
		return err
	}

	followReqSlc, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = f.producer.Produce(followReqSlc, "unfollow")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
