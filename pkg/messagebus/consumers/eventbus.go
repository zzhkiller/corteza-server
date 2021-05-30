package consumers

import (
	"context"

	"github.com/cortezaproject/corteza-server/pkg/eventbus"
	"github.com/cortezaproject/corteza-server/system/service/event"
)

type (
	Dispatcher interface {
		Dispatch(ctx context.Context, ev eventbus.Event)
	}

	EventbusConsumer struct {
		queue      string
		handle     ConsumerType
		dispatcher Dispatcher
	}
)

func NewEventbusConsumer(q string) *EventbusConsumer {
	h := &EventbusConsumer{
		queue:      q,
		handle:     ConsumerEventbus,
		dispatcher: eventbus.Service(),
	}

	return h
}

func (cq *EventbusConsumer) Write(ctx context.Context, p []byte) error {
	cq.dispatcher.Dispatch(ctx, event.QueueOnMessage(cq.queue, p))
	return nil
}
