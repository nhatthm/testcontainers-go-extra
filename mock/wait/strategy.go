package wait

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/testcontainers/testcontainers-go/wait"
)

// StrategyMocker is Strategy mocker.
type StrategyMocker func(tb testing.TB) *Strategy

// NoMockStrategy is no mock Strategy.
var NoMockStrategy = MockStrategy()

var _ wait.Strategy = (*Strategy)(nil)

// Strategy is a wait.Strategy.
type Strategy struct {
	mock.Mock
}

// WaitUntilReady satisfies wait.Strategy interface.
func (s *Strategy) WaitUntilReady(ctx context.Context, target wait.StrategyTarget) error {
	return s.Called(ctx, target).Error(0)
}

// mockStrategy mocks wait.Strategy interface.
func mockStrategy(mocks ...func(s *Strategy)) *Strategy {
	s := &Strategy{}

	for _, m := range mocks {
		m(s)
	}

	return s
}

// MockStrategy creates Strategy mock with cleanup to ensure all the expectations are met.
func MockStrategy(mocks ...func(s *Strategy)) StrategyMocker {
	return func(tb testing.TB) *Strategy {
		tb.Helper()

		s := mockStrategy(mocks...)

		tb.Cleanup(func() {
			assert.True(tb, s.Mock.AssertExpectations(tb))
		})

		return s
	}
}
