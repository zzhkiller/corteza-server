package consumer

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/cortezaproject/corteza-server/pkg/messagebus/store"
	mt "github.com/cortezaproject/corteza-server/pkg/messagebus/types"
	"github.com/cortezaproject/corteza-server/system/types"
	"github.com/stretchr/testify/require"
)

type (
	mockClient struct {
		get     func(context.Context, string) (types.QueueMessageSet, error)
		add     func(context.Context, string, []byte) (err error)
		process func(context.Context, types.QueueMessage) (err error)

		setStore func(store.QueueStorer)
		getStore func() store.QueueStorer
	}
)

var (
	fooMessageSet = types.QueueMessageSet{
		{ID: 1},
		{ID: 2},
	}

	successfulClient = mockClient{
		get: func(c context.Context, q string) (types.QueueMessageSet, error) {
			return fooMessageSet, nil
		},
		add: func(c context.Context, q string, p []byte) error {
			return nil
		},
		process: func(c context.Context, m types.QueueMessage) error {
			return nil
		},
	}

	unsuccessfulClient = mockClient{
		get: func(c context.Context, s string) (types.QueueMessageSet, error) {
			return types.QueueMessageSet{}, errors.New("could not get messages")
		},
		add: func(c context.Context, q string, p []byte) error {
			return errors.New("could not write messages")
		},
		process: func(c context.Context, m types.QueueMessage) error {
			return errors.New("could not process messages")
		},
	}
)

func Test_handlerSqlWrite(t *testing.T) {
	var (
		ctx = context.Background()
		tcc = []struct {
			name    string
			err     error
			payload []byte
			expect  types.QueueMessageSet
			client  store.StoreClient
		}{
			{
				name:   "write success",
				err:    nil,
				expect: fooMessageSet,
				client: &successfulClient,
			},
			{
				name:   "write error",
				err:    errors.New("could not write messages"),
				expect: types.QueueMessageSet{},
				client: &unsuccessfulClient,
			},
		}
	)

	for _, tc := range tcc {
		t.Run(tc.name, func(t *testing.T) {
			req := require.New(t)
			h := StoreConsumer{queue: "foobar", handle: mt.ConsumerStore, client: tc.client}

			err := h.Write(ctx, []byte("foo bar"))

			req.Equal(tc.err, err)
		})
	}
}

func (mc *mockClient) Get(ctx context.Context, q string) (set types.QueueMessageSet, err error) {
	return mc.get(ctx, q)
}

func (mc *mockClient) Add(ctx context.Context, q string, p []byte) error {
	return mc.add(ctx, q, p)
}

func (mc *mockClient) Process(ctx context.Context, m types.QueueMessage) error {
	return mc.process(ctx, m)
}

func (mc *mockClient) GetStore() store.QueueStorer {
	return mc.getStore()
}

func (mc *mockClient) SetStore(s store.QueueStorer) {
	mc.setStore(s)
}

func makeDelay(d time.Duration) *time.Duration {
	return &d
}
