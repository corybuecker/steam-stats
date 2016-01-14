package ratelimiters

import "time"

type ClockInterface interface {
	Sleep(time.Duration)
	Now() time.Time
}

type Clock struct{}

func (clock *Clock) Sleep(d time.Duration) { time.Sleep(d) }
func (clock *Clock) Now() time.Time        { return time.Now() }

type GiantBombRateLimiter struct {
	Clock ClockInterface
}

func (rateLimiter *GiantBombRateLimiter) ObeyRateLimit() error {
	if rateLimiter.Clock == nil {
		rateLimiter.Clock = &Clock{}
	}

	rateLimiter.Clock.Sleep(time.Second * 2)

	return nil
}
