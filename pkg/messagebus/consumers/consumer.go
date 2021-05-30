package consumers

import (
	"context"

	"github.com/cortezaproject/corteza-server/pkg/messagebus"
)

const (
	ConsumerCorteza  ConsumerType = "corteza"
	ConsumerNoop     ConsumerType = "noop"
	ConsumerRedis    ConsumerType = "redis"
	ConsumerStore    ConsumerType = "store"
	ConsumerEventbus ConsumerType = "eventbus"
)

type (
	ConsumerType string

	Consumer interface {
		Writer
		Storer
	}

	Writer interface {
		Write(ctx context.Context, p []byte) error
	}

	Storer interface {
		SetStore(storer messagebus.QueueStorer)
	}

	QueueStorer interface {
		SearchMessagebusQueueMessages(ctx context.Context, f QueueMessageFilter) (QueueMessageSet, QueueMessageFilter, error)
		CreateMessagebusQueueMessage(ctx context.Context, rr ...*QueueMessage) error
		UpdateMessagebusQueueMessage(ctx context.Context, rr ...*QueueMessage) error
	}
)

func ConsumerTypes() []ConsumerType {
	return []ConsumerType{
		ConsumerCorteza,
		ConsumerEventbus,
		ConsumerRedis,
		ConsumerStore,
	}
}
