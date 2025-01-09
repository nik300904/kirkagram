package kafka

import (
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"kirkagram/internal/config"
	"log/slog"
)

type Producer struct {
	producer *kafka.Producer
	log      *slog.Logger
}

var ErrUnknownType = errors.New("unknown type")

func NewProducer(cfg config.Config, log *slog.Logger) *Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.Address,
	})
	if err != nil {
		panic(err)
	}

	return &Producer{
		producer: p,
		log:      log,
	}
}

func (p *Producer) Produce(msg []byte, topic *string) error {
	const op = "kafka.Produce"

	deliveryChan := make(chan kafka.Event)

	if err := p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     topic,
			Partition: kafka.PartitionAny,
		},
		Value: msg,
	}, deliveryChan); err != nil {
		return err
	}

	e := <-deliveryChan
	switch ev := e.(type) {
	case *kafka.Message:
		if ev.TopicPartition.Error != nil {
			return fmt.Errorf("%s %w", op, ev.TopicPartition.Error)
		}
		return nil
	case kafka.Error:
		return fmt.Errorf("%s %w", op, ev)
	default:
		return fmt.Errorf("%s %w", op, ErrUnknownType)
	}
}

func (p *Producer) Close() {
	p.producer.Flush(-1)
	p.producer.Close()
}
