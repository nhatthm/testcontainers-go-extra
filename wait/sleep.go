package wait

import (
	"context"
	"time"

	"github.com/testcontainers/testcontainers-go/wait"
)

var _ wait.Strategy = (*SleepStrategy)(nil)

// SleepStrategy sleeps for an amount of time without checking anything.
type SleepStrategy struct {
	duration time.Duration
}

// WaitUntilReady sleeps for an amount of time. It will return an error if context is canceled.
func (s *SleepStrategy) WaitUntilReady(ctx context.Context, _ wait.StrategyTarget) error {
	select {
	case <-ctx.Done():
		return ctx.Err()

	case <-time.After(s.duration):
		return nil
	}
}

// Sleep will sleep for an amount of time without checking anything.
func Sleep(d time.Duration) *SleepStrategy {
	return &SleepStrategy{duration: d}
}
