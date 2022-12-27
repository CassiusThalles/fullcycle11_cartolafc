package event

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"mygolangapp/pkg/uow"
)

type ProcessEventStrategy interface {
	Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error
}