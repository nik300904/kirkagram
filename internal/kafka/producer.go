package kafka

import (
	"errors"
	"fmt"
	"github.com/IBM/sarama"
	"kirkagram/internal/config"
	"log/slog"
	"strconv"
)

type Producer struct {
	producer sarama.SyncProducer
	log      *slog.Logger
}

var ErrUnknownType = errors.New("unknown type")

func NewProducer(cfg *config.Config, log *slog.Logger) *Producer {
	p, err := sarama.NewSyncProducer([]string{"localhost:29092"}, nil)
	if err != nil {
		panic(err)
	}

	return &Producer{
		producer: p,
		log:      log,
	}
}

func (p *Producer) Produce(msg []byte, topic string) error {
	const op = "kafka.Produce"

	produceMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msg),
	}

	partition, offset, err := p.producer.SendMessage(produceMsg)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	p.log.Info(
		"Message sent",
		slog.String("topic", topic),
		slog.String("partition", strconv.FormatInt(int64(partition), 10)),
		slog.String(
			"offset",
			strconv.FormatInt(int64(offset), 10),
		),
	)

	return nil
}

func (p *Producer) Close() {
	p.producer.Close()
}
