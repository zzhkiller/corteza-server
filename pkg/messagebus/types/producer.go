package types

import (
	"context"
	"time"

	"github.com/cortezaproject/corteza-server/pkg/messagebus/store"
	"github.com/cortezaproject/corteza-server/system/types"
)

type (
	Producer interface {
		Reader
		Subscriber
		Poller
		store.Storer
	}

	Subscriber interface {
		Subscribe(ctx context.Context) <-chan interface{}
	}

	Poller interface {
		Poll(ctx context.Context) <-chan time.Time
	}

	Reader interface {
		Read(ctx context.Context) (types.QueueMessageSet, error)
	}
)
