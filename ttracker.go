package ttracker

import (
	"time"
)

type (
	TimeoutTracker struct {
		Handler LabelHandler
		labels  map[string]time.Duration
		name    string

		started time.Time
		timeout time.Duration
	}

	LabelHandler func(name, label string, timeout time.Duration)
)

var (
	DefaultHandler LabelHandler
)

func NewTimeoutTracker(name string, timeout time.Duration) *TimeoutTracker {
	return &TimeoutTracker{
		Handler: DefaultHandler,

		name:    name,
		timeout: timeout,
	}
}

func Start(name string, timeout time.Duration) *TimeoutTracker {
	return NewTimeoutTracker(name, timeout).Start()
}

func (t *TimeoutTracker) Track(label string) {
	t.labels[label] = time.Since(t.started)
}

func (t *TimeoutTracker) Start() *TimeoutTracker {
	t.started = time.Now()
	t.labels = map[string]time.Duration{}
	return t
}

func (t *TimeoutTracker) Stop() {
	timeout := time.Since(t.started)
	if timeout >= t.timeout {
		if t.Handler != nil {
			for l, tm := range t.labels {
				t.Handler(t.name, l, tm)
			}
		}
	}
}
