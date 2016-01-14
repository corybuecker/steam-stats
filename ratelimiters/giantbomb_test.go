package ratelimiters

import (
	"testing"
	"time"
)

type TestClock struct {
	sleepDuration time.Duration
}

func (clock *TestClock) Sleep(d time.Duration) {
	clock.sleepDuration = d
}
func (clock *TestClock) Now() time.Time {
	return time.Date(2009, time.January, 1, 0, 0, 0, 0, time.UTC)
}

func TestCreatedClock(t *testing.T) {
	rateLimiter := GiantBombRateLimiter{}
	rateLimiter.ObeyRateLimit()
	if rateLimiter.Clock == nil {
		t.Fatalf("should have created a clock")
	}
}

func TestCorrectSleep(t *testing.T) {

	rateLimiter := GiantBombRateLimiter{Clock: &TestClock{}}

	rateLimiter.ObeyRateLimit()
	if rateLimiter.Clock.(*TestClock).sleepDuration != time.Second*2 {
		t.Fatalf("expected around one second, got %s", rateLimiter.Clock.(*TestClock).sleepDuration)
	}
}
