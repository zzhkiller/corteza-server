package types

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_settingsUnmarshal(t *testing.T) {
	var (
		makeDelay = func(s time.Duration) *time.Duration {
			return &s
		}

		tcc = []struct {
			name    string
			payload string
			err     error
			expect  QueueMeta
		}{
			{
				name:    "settings defaults",
				payload: `{}`,
				err:     nil,
				expect:  QueueMeta{PollDelay: nil, DispatchEvents: false},
			},
			{
				name:    "settings enabled dispatch events",
				payload: `{"poll_delay": "7s", "dispatch_events": true}`,
				err:     nil,
				expect:  QueueMeta{PollDelay: makeDelay(time.Second * 7), DispatchEvents: true},
			},
			{
				name:    "settings disabled dispatch events",
				payload: `{"poll_delay": "7s", "dispatch_events": false}`,
				err:     nil,
				expect:  QueueMeta{PollDelay: makeDelay(time.Second * 7), DispatchEvents: false},
			},
			{
				name:    "settings invalid poll delay",
				payload: `{"poll_delay": "7seconds", "dispatch_events": false}`,
				err:     nil,
				expect:  QueueMeta{PollDelay: nil, DispatchEvents: false},
			},
		}
	)

	for _, tc := range tcc {
		t.Run(tc.name, func(t *testing.T) {
			req := require.New(t)

			s := &QueueMeta{}
			err := json.Unmarshal([]byte(tc.payload), s)

			req.Equal(tc.err, err)
			req.Equal(tc.expect, *s)
		})
	}
}
