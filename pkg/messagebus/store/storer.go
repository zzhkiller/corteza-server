package store

import (
	"context"

	"github.com/cortezaproject/corteza-server/system/types"
)

type (
	Storer interface {
		SetStore(QueueStorer)
		GetStore() QueueStorer
	}

	QueueStorer interface {
		SearchQueues(ctx context.Context, f types.QueueFilter) (types.QueueSet, types.QueueFilter, error)
		SearchQueueMessages(ctx context.Context, f types.QueueMessageFilter) (types.QueueMessageSet, types.QueueMessageFilter, error)
		CreateQueueMessage(ctx context.Context, rr ...*types.QueueMessage) error
		UpdateQueueMessage(ctx context.Context, rr ...*types.QueueMessage) error
	}

	// Store struct {
	// 	s QueueStorer
	// }

	// QueueMessageFilter struct {
	// 	Queue string

	// 	Processed filter.State `json:"processed"`

	// 	filter.Sorting
	// 	filter.Paging
	// }

	// QueueFilter struct {
	// 	Queue    string `json:"queue"`
	// 	Consumer string `json:"handler"`

	// 	Deleted filter.State `json:"deleted"`

	// 	// Check fn is called by store backend for each resource found function can
	// 	// modify the resource and return false if store should not return it
	// 	//
	// 	// Store then loads additional resources to satisfy the paging parameters
	// 	// Check func(*Queue) (bool, error) `json:"-"`

	// 	filter.Sorting
	// 	filter.Paging
	// }
)

// func NewStore(s QueueStorer) *Store {
// 	return &Store{s: s}
// }

// func (s *Store) SearchQueues(ctx context.Context, f QueueFilter) (set Queue, filter QueueFilter, err error) {
// 	return
// }

// func (s *Store) SearchQueueMessages(ctx context.Context, f QueueMessageFilter) (set QueueMessageSet, filter QueueMessageFilter, err error) {
// 	return
// }

// func (s *Store) CreateQueueMessage(ctx context.Context, f QueueMessageFilter) (err error) {
// 	return
// }
