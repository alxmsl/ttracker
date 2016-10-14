package ttracker

import (
	"time"
)

type (
	Timeout struct {
		Label    string
		Duration time.Duration
		Elapsed  time.Duration
	}

	TimeoutTracker struct {
		Duration time.Duration
		Handler  LabelHandler
		Name     string
		Timeouts []Timeout

		last      time.Time
		started   time.Time
		threshold time.Duration
	}

	LabelHandler func(tracker *TimeoutTracker)
)

var (
	DefaultHandler LabelHandler
)

func NewTimeoutTracker(name string, threshold time.Duration) *TimeoutTracker {
	return &TimeoutTracker{
		Handler:  DefaultHandler,
		Name:     name,
		Timeouts: []Timeout{},

		threshold: threshold,
	}
}

func Start(name string, threshold time.Duration) *TimeoutTracker {
	return NewTimeoutTracker(name, threshold).Start()
}

func (t *TimeoutTracker) Track(label string) {
	now := time.Now()
	t.Timeouts = append(t.Timeouts, Timeout{
		Label:    label,
		Duration: now.Sub(t.last),
		Elapsed:  now.Sub(t.started),
	})
	t.last = now
}

func (t *TimeoutTracker) Start() *TimeoutTracker {
	t.last = time.Now()
	t.started = t.last
	return t
}

func (t *TimeoutTracker) Stop() {
	t.Duration = time.Since(t.started)
	if t.Duration >= t.threshold && t.Handler != nil {
		t.Handler(t)
	}
}
