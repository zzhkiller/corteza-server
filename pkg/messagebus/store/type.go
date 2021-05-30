package store

import (
	"github.com/cortezaproject/corteza-server/pkg/messagebus/consumers"
)

type (
	message struct {
		q string
		p []byte
	}

	Queue struct {
		consumer consumers.Consumer

		// settings are used in messagebus for handling
		// throttling, polling settings
		settings *Settings
	}

	Settings struct {
		Name     string
		Consumer consumers.Consumer
	}

	QueueSet map[string]*Queue
)
