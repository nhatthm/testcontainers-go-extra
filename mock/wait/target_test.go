package wait_test

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/assert"

	"go.nhat.io/testcontainers-go-extra/mock/wait"
)

func TestStrategyTarget_Host(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mock           wait.StrategyTargetMocker
		expectedResult string
		expectedError  error
	}{
		{
			scenario: "error",
			mock: wait.MockStrategyTarget(func(t *wait.StrategyTarget) {
				t.On("Host", context.Background()).
					Return("", errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "error",
			mock: wait.MockStrategyTarget(func(t *wait.StrategyTarget) {
				t.On("Host", context.Background()).
					Return("localhost", nil)
			}),
			expectedResult: "localhost",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			result, err := tc.mock(t).Host(context.Background())

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestStrategyTarget_MappedPort(t *testing.T) {
	t.Parallel()

	port := nat.Port("8080")

	testCases := []struct {
		scenario       string
		mock           wait.StrategyTargetMocker
		expectedResult nat.Port
		expectedError  error
	}{
		{
			scenario: "string and error",
			mock: wait.MockStrategyTarget(func(t *wait.StrategyTarget) {
				t.On("MappedPort", context.Background(), port).
					Return("", errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "string and no error",
			mock: wait.MockStrategyTarget(func(t *wait.StrategyTarget) {
				t.On("MappedPort", context.Background(), port).
					Return("8080", nil)
			}),
			expectedResult: port,
		},
		{
			scenario: "port and error",
			mock: wait.MockStrategyTarget(func(t *wait.StrategyTarget) {
				t.On("MappedPort", context.Background(), port).
					Return(nat.Port(""), errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "port and no error",
			mock: wait.MockStrategyTarget(func(t *wait.StrategyTarget) {
				t.On("MappedPort", context.Background(), port).
					Return(port, nil)
			}),
			expectedResult: port,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			result, err := tc.mock(t).MappedPort(context.Background(), port)

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestStrategyTarget_Logs(t *testing.T) {
	t.Parallel()

	rc := io.NopCloser(nil)

	testCases := []struct {
		scenario       string
		mock           wait.StrategyTargetMocker
		expectedResult io.ReadCloser
		expectedError  error
	}{
		{
			scenario: "error",
			mock: wait.MockStrategyTarget(func(t *wait.StrategyTarget) {
				t.On("Logs", context.Background()).
					Return(nil, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "no error",
			mock: wait.MockStrategyTarget(func(t *wait.StrategyTarget) {
				t.On("Logs", context.Background()).
					Return(rc, nil)
			}),
			expectedResult: rc,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			result, err := tc.mock(t).Logs(context.Background())

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestStrategyTarget_Exec(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mock           wait.StrategyTargetMocker
		expectedResult int
		expectedError  error
	}{
		{
			scenario: "error",
			mock: wait.MockStrategyTarget(func(t *wait.StrategyTarget) {
				t.On("Exec", context.Background(), []string{"test"}).
					Return(1, errors.New("error"))
			}),
			expectedResult: 1,
			expectedError:  errors.New("error"),
		},
		{
			scenario: "no error",
			mock: wait.MockStrategyTarget(func(t *wait.StrategyTarget) {
				t.On("Exec", context.Background(), []string{"test"}).
					Return(0, nil)
			}),
			expectedResult: 0,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			result, err := tc.mock(t).Exec(context.Background(), []string{"test"})

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestStrategyTarget_State(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mock           wait.StrategyTargetMocker
		expectedResult *types.ContainerState
		expectedError  error
	}{
		{
			scenario: "error",
			mock: wait.MockStrategyTarget(func(t *wait.StrategyTarget) {
				t.On("State", context.Background()).
					Return(nil, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "no error",
			mock: wait.MockStrategyTarget(func(t *wait.StrategyTarget) {
				t.On("State", context.Background()).
					Return(&types.ContainerState{
						Status:  "running",
						Running: true,
					}, nil)
			}),
			expectedResult: &types.ContainerState{
				Status:  "running",
				Running: true,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			result, err := tc.mock(t).State(context.Background())

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
