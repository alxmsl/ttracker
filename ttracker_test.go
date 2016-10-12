package ttracker

import (
	"strconv"
	"testing"
	"time"

	. "gopkg.in/check.v1"
)

type (
	TrackerSuite struct {
		tracker *TimeoutTracker
	}
)

var (
	handled      bool
	handledCount int
	handler      = func(name, label string, timeout time.Duration) {
		handled = true
		handledCount += 1
	}
	tracker = NewTimeoutTracker("tracker", 10*time.Millisecond)
)

func init() {
	tracker.Handler = handler
	_ = Suite(&TrackerSuite{
		tracker: tracker,
	})
}

func Test(t *testing.T) { TestingT(t) }

func (s *TrackerSuite) SetUpTest(c *C) {
	handled = false
	handledCount = 0
}

func (s *TrackerSuite) TestStopAndDontHandle(c *C) {
	s.tracker.Start()
	s.tracker.Stop()
	c.Assert(handled, Equals, false)
}

func (s *TrackerSuite) TestStopAndDontHandleBecauseNoTracks(c *C) {
	s.tracker.Start()
	<-time.NewTimer(11 * time.Millisecond).C
	s.tracker.Stop()
	c.Assert(handled, Equals, false)
}

func (s *TrackerSuite) TestStopAndDontHandleOneTime(c *C) {
	s.tracker.Start()
	s.tracker.Track("a")
	<-time.NewTimer(11 * time.Millisecond).C
	s.tracker.Stop()
	c.Assert(handled, Equals, true)
	c.Assert(handledCount, Equals, 1)
}

func (s *TrackerSuite) TestStopAndHandle10Times(c *C) {
	s.tracker.Start()
	for i := 0; i < 10; i += 1 {
		s.tracker.Track(strconv.Itoa(i))
	}
	<-time.NewTimer(11 * time.Millisecond).C
	s.tracker.Stop()
	c.Assert(handled, Equals, true)
	c.Assert(handledCount, Equals, 10)
}

func (s *TrackerSuite) BenchmarkTracking(c *C) {
	predefinedLabels := make([]string, c.N)
	for i := 0; i < cap(predefinedLabels); i += 1 {
		predefinedLabels[i] = strconv.Itoa(i)
	}

	s.tracker.Start()
	c.ResetTimer()
	for i := 0; i < c.N; i++ {
		s.tracker.Track(predefinedLabels[i])
	}
}
