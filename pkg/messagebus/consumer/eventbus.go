package consumer

import (
	"context"

	"github.com/cortezaproject/corteza-server/pkg/eventbus"
	"github.com/cortezaproject/corteza-server/pkg/messagebus/types"
	"github.com/cortezaproject/corteza-server/system/service/event"
	st "github.com/cortezaproject/corteza-server/system/types"
)

type (
	Dispatcher interface {
		Dispatch(ctx context.Context, ev eventbus.Event)
	}

	EventbusConsumer struct {
		queue      string
		handle     types.ConsumerType
		dispatcher Dispatcher
	}
)

func NewEventbusConsumer(q string) *EventbusConsumer {
	h := &EventbusConsumer{
		queue:      q,
		handle:     types.ConsumerEventbus,
		dispatcher: eventbus.Service(),
	}

	return h
}

func (cq *EventbusConsumer) Write(ctx context.Context, p []byte) error {
	cq.dispatcher.Dispatch(ctx, event.QueueOnMessage(makeEvent(cq.queue, p)))
	return nil
}

func makeEvent(q string, p []byte) *st.QueueMessage {
	return &st.QueueMessage{
		Queue:   q,
		Payload: p,
	}
}
