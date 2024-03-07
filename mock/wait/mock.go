package wait

import (
	"testing"

	"github.com/testcontainers/testcontainers-go/wait"
)

// StrategyMocker is Strategy mocker.
type StrategyMocker func(tb testing.TB) *Strategy

// NopStrategy is no mock Strategy.
var NopStrategy = MockStrategy()

var _ wait.Strategy = (*Strategy)(nil)

// MockStrategy creates Strategy mock with cleanup to ensure all the expectations are met.
func MockStrategy(mocks ...func(s *Strategy)) StrategyMocker {
	return func(tb testing.TB) *Strategy {
		tb.Helper()

		s := NewStrategy(tb)

		for _, m := range mocks {
			m(s)
		}

		return s
	}
}

// StrategyTargetMocker is StrategyTarget mocker.
type StrategyTargetMocker func(tb testing.TB) *StrategyTarget

// NopStrategyTarget is no mock StrategyTarget.
var NopStrategyTarget = MockStrategyTarget()

var _ wait.StrategyTarget = (*StrategyTarget)(nil)

// MockStrategyTarget creates StrategyTarget mock with cleanup to ensure all the expectations are met.
func MockStrategyTarget(mocks ...func(t *StrategyTarget)) StrategyTargetMocker {
	return func(tb testing.TB) *StrategyTarget {
		tb.Helper()

		t := NewStrategyTarget(tb)

		for _, m := range mocks {
			m(t)
		}

		return t
	}
}
