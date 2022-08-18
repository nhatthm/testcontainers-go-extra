package wait_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/testcontainers-extra/wait"
)

func TestSleep_Background(t *testing.T) {
	t.Parallel()

	s := wait.Sleep(50 * time.Millisecond)
	err := s.WaitUntilReady(context.Background(), nil)

	assert.NoError(t, err)
}

func TestSleep_Canceled(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	cancel()

	s := wait.Sleep(60 * time.Millisecond)
	err := s.WaitUntilReady(ctx, nil)

	assert.ErrorIs(t, err, context.Canceled)
}

func TestSleep_DeadlineExceeded(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	s := wait.Sleep(60 * time.Millisecond)
	err := s.WaitUntilReady(ctx, nil)

	assert.ErrorIs(t, err, context.DeadlineExceeded)
}

func TestSleep_NoTimeout(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	s := wait.Sleep(40 * time.Millisecond)
	err := s.WaitUntilReady(ctx, nil)

	assert.NoError(t, err)
}
