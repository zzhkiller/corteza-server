package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/cast"
)

type (
	Message struct {
		Q string
		P []byte
	}

	Queue struct {
		Name     string
		Consumer Consumer
		Meta     QueueMeta
	}

	QueueSet map[string]*Queue

	// Queue1 struct {
	// 	ID       uint64    `json:"queueID,string"`
	// 	Consumer string    `json:"consumer"`
	// 	Queue    string    `json:"queue"`
	// 	Meta     QueueMeta `json:"meta"`

	// 	CreatedAt time.Time  `json:"createdAt,omitempty"`
	// 	CreatedBy uint64     `json:"createdBy,string" `
	// 	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// 	UpdatedBy uint64     `json:"updatedBy,string,omitempty" `
	// 	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// 	DeletedBy uint64     `json:"deletedBy,string,omitempty" `
	// }

	// // QueueSet []Queue1

	QueueMeta struct {
		PollDelay      *time.Duration `json:"poll_delay"`
		DispatchEvents bool           `json:"dispatch_events"`
	}

	QueueMessage struct {
		Queue   string `json:"queue"`
		Payload string `json:"payload"`
	}

	QueueMessageSet []QueueMessage
)

func (h *QueueMeta) UnmarshalJSON(s []byte) error {
	type Alias QueueMeta

	aux := &struct {
		PollDelay string `json:"poll_delay"`
		*Alias
	}{
		Alias: (*Alias)(h),
	}

	// set default
	h.DispatchEvents = false

	if err := json.Unmarshal(s, aux); err != nil {
		return err
	}

	if d, err := cast.ToDurationE(aux.PollDelay); err == nil {
		h.PollDelay = &d
	}

	return nil
}

func (m QueueMeta) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m QueueMeta) MarshalJSON() ([]byte, error) {

	pollDelay := ""
	if m.PollDelay != nil {
		pollDelay = m.PollDelay.String()
	}

	return json.Marshal(struct {
		PollDelay      string `json:"poll_delay"`
		DispatchEvents bool   `json:"dispatch_events"`
	}{
		PollDelay:      pollDelay,
		DispatchEvents: m.DispatchEvents,
	})
}

func (m *QueueMeta) Scan(value interface{}) error {
	switch value.(type) {
	case nil:
		*m = QueueMeta{}
	case []uint8:
		if err := json.Unmarshal(value.([]byte), m); err != nil {
			return fmt.Errorf("cannot scan '%v' into QueueMeta", value)
		}
	}

	return nil
}

func (s *Queue) CanDispatch() bool {
	return s.Meta.DispatchEvents
}
