package store

import (
	"context"
	"time"

	"github.com/cortezaproject/corteza-server/pkg/filter"
	"github.com/cortezaproject/corteza-server/pkg/id"
	"github.com/cortezaproject/corteza-server/system/types"
)

type (
	Client interface {
		Add(ctx context.Context, q string, p []byte) (err error)
		Get(ctx context.Context, q string) (list types.QueueMessageSet, err error)
		Process(ctx context.Context, m types.QueueMessage) (err error)
	}

	StoreClient interface {
		Client
		Storer
	}

	sClient struct {
		store QueueStorer
	}
)

func NewClient(s QueueStorer) *sClient {
	return &sClient{store: s}
}

func (c *sClient) Process(ctx context.Context, m types.QueueMessage) (err error) {
	err = c.store.UpdateQueueMessage(ctx, &types.QueueMessage{
		ID:        m.ID,
		Queue:     m.Queue,
		Payload:   m.Payload,
		Created:   m.Created,
		Processed: now(),
	})

	return
}

func (c *sClient) Add(ctx context.Context, q string, payload []byte) (err error) {
	err = c.store.CreateQueueMessage(ctx, &types.QueueMessage{
		ID:      nextID(),
		Queue:   q,
		Created: now(),
		Payload: payload,
	})

	return
}

func (c *sClient) Get(ctx context.Context, q string) (list types.QueueMessageSet, err error) {
	list, _, err = c.store.SearchQueueMessages(ctx, types.QueueMessageFilter{
		Queue:     q,
		Processed: filter.StateExcluded})

	return
}

func (c *sClient) GetStore() QueueStorer {
	return c.store
}

func (c *sClient) SetStore(s QueueStorer) {
	c.store = s
}

func nextID() uint64 {
	return id.Next()
}

func now() *time.Time {
	t := time.Now()
	return &t
}
