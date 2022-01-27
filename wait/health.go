package wait

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	defaultStartPeriod  = time.Duration(0)
	defaultRetries      = 3
	defaultTestTimeout  = 10 * time.Second
	defaultTestInterval = 5 * time.Second
)

var _ wait.Strategy = (*HealthCheckStrategy)(nil)

// HealthCheckStrategy is a strategy for doing health check.
type HealthCheckStrategy struct {
	test         HealthCheckTest
	testInterval time.Duration
	testTimeout  time.Duration
	retries      int
	startPeriod  time.Duration
}

// WithTestInterval sets the interval between retries.
func (s *HealthCheckStrategy) WithTestInterval(interval time.Duration) *HealthCheckStrategy {
	s.testInterval = interval

	return s
}

// WithTestTimeout sets timeout for running the test.
func (s *HealthCheckStrategy) WithTestTimeout(timeout time.Duration) *HealthCheckStrategy {
	s.testTimeout = timeout

	return s
}

// WithRetries sets a number of retries for running the test.
func (s *HealthCheckStrategy) WithRetries(retries int) *HealthCheckStrategy {
	s.retries = retries

	return s
}

// WithStartPeriod sets a period when retries are not counted. This is helpful when a container takes some time to spin
// up before it's ready for health checking.
func (s *HealthCheckStrategy) WithStartPeriod(duration time.Duration) *HealthCheckStrategy {
	s.startPeriod = duration

	return s
}

func (s *HealthCheckStrategy) testTarget(target wait.StrategyTarget) (success bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.testTimeout)
	defer cancel()

	success, err = s.test.Test(ctx, target)
	if errors.Is(err, context.DeadlineExceeded) {
		return false, nil
	}

	return
}

// WaitUntilReady runs the health check test.
func (s *HealthCheckStrategy) WaitUntilReady(ctx context.Context, target wait.StrategyTarget) error {
	pollInternal := time.Duration(0)
	retry := 0

	startTime := time.Now()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case <-time.After(pollInternal):
			success, err := s.testTarget(target) // nolint: contextcheck
			if err != nil {
				return err
			}

			if success {
				return nil
			}

			elapsedTime := time.Since(startTime)

			if elapsedTime > s.startPeriod {
				retry++
			}

			if retry > s.retries {
				return fmt.Errorf("health check failed: %w", ErrMaxRetriesExceeded)
			}

			pollInternal = s.testInterval
		}
	}
}

// HealthCheckTest tests for health check.
type HealthCheckTest interface {
	Test(ctx context.Context, target StrategyTarget) (success bool, err error)
}

// HealthCheckTestFunc is an inline test for health check.
type HealthCheckTestFunc func(ctx context.Context, target StrategyTarget) (success bool, err error)

// Test checks if a container is healthy.
func (f HealthCheckTestFunc) Test(ctx context.Context, target StrategyTarget) (success bool, err error) {
	return f(ctx, target)
}

// ForHealthCheck creates a new health check with default arguments.
func ForHealthCheck(f HealthCheckTestFunc) *HealthCheckStrategy {
	return &HealthCheckStrategy{
		test:         f,
		retries:      defaultRetries,
		testTimeout:  defaultTestTimeout,
		testInterval: defaultTestInterval,
		startPeriod:  defaultStartPeriod,
	}
}
