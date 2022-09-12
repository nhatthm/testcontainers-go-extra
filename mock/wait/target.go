package wait

import (
	"context"
	"io"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/testcontainers/testcontainers-go/wait"
)

// StrategyTargetMocker is StrategyTarget mocker.
type StrategyTargetMocker func(tb testing.TB) *StrategyTarget

// NoMockStrategyTarget is no mock StrategyTarget.
var NoMockStrategyTarget = MockStrategyTarget()

var _ wait.StrategyTarget = (*StrategyTarget)(nil)

// StrategyTarget is a wait.StrategyTarget.
type StrategyTarget struct {
	mock.Mock
}

// Host satisfies wait.StrategyTarget interface.
func (s *StrategyTarget) Host(ctx context.Context) (string, error) {
	result := s.Called(ctx)

	return result.String(0), result.Error(1)
}

// MappedPort satisfies wait.StrategyTarget interface.
func (s *StrategyTarget) MappedPort(ctx context.Context, port nat.Port) (nat.Port, error) {
	result := s.Called(ctx, port)

	p := result.Get(0)
	err := result.Error(1)

	if s, ok := p.(string); ok {
		return nat.Port(s), err
	}

	return p.(nat.Port), err
}

// Logs satisfies wait.StrategyTarget interface.
func (s *StrategyTarget) Logs(ctx context.Context) (io.ReadCloser, error) {
	result := s.Called(ctx)

	rc := result.Get(0)
	err := result.Error(1)

	if rc == nil {
		return nil, err
	}

	return rc.(io.ReadCloser), err
}

// Exec satisfies wait.StrategyTarget interface.
func (s *StrategyTarget) Exec(ctx context.Context, cmd []string) (int, io.Reader, error) {
	result := s.Called(ctx, cmd)

	r1, r2, r3 := result.Int(0), result.Get(1), result.Error(2)

	if r2 == nil {
		return r1, nil, r3
	}

	return r1, r2.(io.Reader), r3
}

// Ports satisfies wait.StrategyTarget interface.
func (s *StrategyTarget) Ports(ctx context.Context) (nat.PortMap, error) {
	result := s.Called(ctx)

	r1, r2 := result.Get(0), result.Error(1)

	if r1 == nil {
		return nil, r2
	}

	return r1.(nat.PortMap), r2
}

// State satisfies wait.StrategyTarget interface.
func (s *StrategyTarget) State(ctx context.Context) (*types.ContainerState, error) {
	result := s.Called(ctx)

	state := result.Get(0)
	err := result.Error(1)

	if state == nil {
		return nil, err
	}

	return state.(*types.ContainerState), err
}

// mockStrategyTarget mocks wait.StrategyTarget interface.
func mockStrategyTarget(mocks ...func(t *StrategyTarget)) *StrategyTarget {
	t := &StrategyTarget{}

	for _, m := range mocks {
		m(t)
	}

	return t
}

// MockStrategyTarget creates StrategyTarget mock with cleanup to ensure all the expectations are met.
func MockStrategyTarget(mocks ...func(t *StrategyTarget)) StrategyTargetMocker {
	return func(tb testing.TB) *StrategyTarget {
		tb.Helper()

		t := mockStrategyTarget(mocks...)

		tb.Cleanup(func() {
			assert.True(tb, t.Mock.AssertExpectations(tb))
		})

		return t
	}
}
