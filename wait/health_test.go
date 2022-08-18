package wait_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/testcontainers-extra/wait"
)

func TestHealthCheckStrategy_WithTestInterval_TestError(t *testing.T) {
	t.Parallel()

	expectedError := errors.New("test error")
	expectedCalled := 1

	called := 0

	s := wait.ForHealthCheck(func(context.Context, wait.StrategyTarget) (success bool, err error) {
		called++

		return false, expectedError
	})

	actual := s.WaitUntilReady(context.Background(), nil)

	assert.Equal(t, expectedError, actual)
	assert.Equal(t, expectedCalled, called)
}

// nolint: paralleltest
func TestHealthCheckStrategy_WithTestInterval_ContextCanceledAfterFirstRetry(t *testing.T) {
	timeout := 15 * time.Millisecond
	called := 0

	s := wait.ForHealthCheck(func(context.Context, wait.StrategyTarget) (success bool, err error) {
		called++

		return false, nil
	}).
		WithTestTimeout(10 * time.Millisecond).
		WithTestInterval(10 * time.Millisecond).
		WithStartPeriod(time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	startTime := time.Now()
	err := s.WaitUntilReady(ctx, nil)
	elapsedTime := time.Since(startTime)

	assert.ErrorIs(t, err, context.DeadlineExceeded)

	expectedTime := 15 * time.Millisecond
	assertInDeltaDurationf(t, elapsedTime, expectedTime, 5*time.Millisecond, "strategy should fail within %s because of the given context, was %s", expectedTime, elapsedTime)

	expectedCalled := 2
	assert.Equal(t, expectedCalled, called, "test was called %d time(s), expected %d", called, expectedCalled)
}

// nolint: paralleltest
func TestHealthCheckStrategy_WithTestInterval_SuccessFirstTry(t *testing.T) {
	called := 0

	s := wait.ForHealthCheck(func(context.Context, wait.StrategyTarget) (success bool, err error) {
		called++

		return true, nil
	}).
		WithTestTimeout(10 * time.Millisecond).
		WithTestInterval(10 * time.Millisecond).
		WithStartPeriod(time.Second)

	startTime := time.Now()
	err := s.WaitUntilReady(context.Background(), nil)
	elapsedTime := time.Since(startTime)

	assert.NoError(t, err)

	expectedTime := time.Duration(0)
	assertInDeltaDurationf(t, elapsedTime, expectedTime, 5*time.Millisecond, "strategy should succeed within %s, was %s", expectedTime, elapsedTime)

	expectedCalled := 1
	assert.Equal(t, expectedCalled, called, "test was called %d time(s), expected %d", called, expectedCalled)
}

// nolint: paralleltest
func TestHealthCheckStrategy_WithTestInterval_SuccessSecondTry(t *testing.T) {
	called := 0

	s := wait.ForHealthCheck(func(context.Context, wait.StrategyTarget) (success bool, err error) {
		called++

		if called != 2 {
			return false, nil
		}

		return true, nil
	}).
		WithTestTimeout(10 * time.Millisecond).
		WithTestInterval(10 * time.Millisecond).
		WithStartPeriod(time.Second)

	startTime := time.Now()
	err := s.WaitUntilReady(context.Background(), nil)
	elapsedTime := time.Since(startTime)

	assert.NoError(t, err)

	expectedTime := 10 * time.Millisecond
	assertInDeltaDurationf(t, elapsedTime, expectedTime, 5*time.Millisecond, "strategy should succeed within %s, was %s", expectedTime, elapsedTime)

	expectedCalled := 2
	assert.Equal(t, expectedCalled, called, "test was called %d time(s), expected %d", called, expectedCalled)
}

// nolint: paralleltest
func TestHealthCheckStrategy_WithTestInterval_FailInStartPeriodNotCountedAsRetry(t *testing.T) {
	// First call takes 0ms.
	// Next 3 calls are not counted because of the start period (35ms).
	// The test interval is 10ms, so the 5th call at 40ms will be counted as a retry.
	// The 6th call is a success
	// Therefore:
	// 	- Expected Time is 50ms
	// 	- Expected Calls is 6
	expectedTime := 50 * time.Millisecond
	expectedCalled := 6

	called := 0

	s := wait.ForHealthCheck(func(context.Context, wait.StrategyTarget) (success bool, err error) {
		called++

		if called != expectedCalled {
			return false, nil
		}

		return true, nil
	}).
		WithTestTimeout(10 * time.Millisecond).
		WithTestInterval(10 * time.Millisecond).
		WithRetries(1).
		WithStartPeriod(35 * time.Millisecond)

	startTime := time.Now()
	err := s.WaitUntilReady(context.Background(), nil)
	elapsedTime := time.Since(startTime)

	assert.NoError(t, err)
	assertInDeltaDurationf(t, elapsedTime, expectedTime, 5*time.Millisecond, "strategy should succeed within %s", expectedTime)
	assert.Equal(t, expectedCalled, called, "test was called %d time(s), expected %d", called, expectedCalled)
}

// nolint: paralleltest
func TestHealthCheckStrategy_WithTestInterval_TestTimeout(t *testing.T) {
	expectedTime := 20 * time.Millisecond
	expectedCalled := 2

	called := 0

	s := wait.ForHealthCheck(func(ctx context.Context, _ wait.StrategyTarget) (success bool, err error) {
		called++

		if called != expectedCalled {
			<-ctx.Done()

			return false, fmt.Errorf("error: %w", ctx.Err())
		}

		return true, nil
	}).
		WithTestTimeout(10 * time.Millisecond).
		WithTestInterval(10 * time.Millisecond).
		WithRetries(1)

	startTime := time.Now()
	err := s.WaitUntilReady(context.Background(), nil)
	elapsedTime := time.Since(startTime)

	assert.NoError(t, err)
	assertInDeltaDurationf(t, elapsedTime, expectedTime, 5*time.Millisecond, "strategy should succeed within %s", expectedTime)
	assert.Equal(t, expectedCalled, called, "test was called %d time(s), expected %d", called, expectedCalled)
}

// nolint: paralleltest
func TestHealthCheckStrategy_WithTestInterval_MaxRetriesExceeded(t *testing.T) {
	s := wait.ForHealthCheck(func(context.Context, wait.StrategyTarget) (success bool, err error) {
		return false, nil
	}).
		WithTestTimeout(10 * time.Millisecond).
		WithTestInterval(10 * time.Millisecond).
		WithRetries(1)

	err := s.WaitUntilReady(context.Background(), nil)
	expected := "health check failed: max retries exceeded"

	assert.ErrorIs(t, err, wait.ErrMaxRetriesExceeded)
	assert.EqualError(t, err, expected)
}

// nolint: unparam
func assertInDeltaDurationf(t assert.TestingT, expected, actual, delta time.Duration, msg string, args ...interface{}) bool {
	return assert.InDeltaf(t, expected, actual, float64(delta), msg, args...)
}
