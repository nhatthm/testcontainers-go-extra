package wait_test

import (
	"context"
	"errors"
	"io"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	waitmock "go.nhat.io/testcontainers-extra/mock/wait"
	"go.nhat.io/testcontainers-extra/wait"
)

func TestForHealthCheckCmd(t *testing.T) {
	t.Parallel()

	pausedState := &types.ContainerState{Status: "paused"}
	restartingState := &types.ContainerState{Status: "restarting"}
	runningState := &types.ContainerState{Status: "running", Running: true}

	testCases := []struct {
		scenario        string
		mockTarget      waitmock.StrategyTargetMocker
		expectedSuccess bool
		expectedError   string
	}{
		{
			scenario: "unable to get state",
			mockTarget: waitmock.MockStrategyTarget(func(t *waitmock.StrategyTarget) {
				t.On("State", isContext).
					Return(nil, errors.New("get state error"))
			}),
			expectedSuccess: false,
			expectedError:   `get state error`,
		},
		{
			scenario: "untestable and could not get logs",
			mockTarget: waitmock.MockStrategyTarget(func(t *waitmock.StrategyTarget) {
				t.On("State", isContext).
					Return(pausedState, nil)

				t.On("Logs", isContext).
					Return(nil, errors.New("get logs error"))
			}),
			expectedSuccess: false,
			expectedError:   `container is paused and unable to get logs: get logs error`,
		},
		{
			scenario: "untestable and could not read logs",
			mockTarget: waitmock.MockStrategyTarget(func(t *waitmock.StrategyTarget) {
				t.On("State", isContext).
					Return(pausedState, nil)

				t.On("Logs", isContext).
					Return(errorReadCloser(errors.New("read logs error")), nil)
			}),
			expectedSuccess: false,
			expectedError:   `container is paused and unable to read logs: read logs error`,
		},
		{
			scenario: "untestable and could read logs",
			mockTarget: waitmock.MockStrategyTarget(func(t *waitmock.StrategyTarget) {
				t.On("State", isContext).
					Return(pausedState, nil)

				t.On("Logs", isContext).
					Return(stringReadCloser("log message"), nil)
			}),
			expectedSuccess: false,
			expectedError:   "container is paused, logs:\nlog message",
		},
		{
			scenario: "untestable and no logs",
			mockTarget: waitmock.MockStrategyTarget(func(t *waitmock.StrategyTarget) {
				t.On("State", isContext).
					Return(pausedState, nil)

				t.On("Logs", isContext).
					Return(nil, nil)
			}),
			expectedSuccess: false,
			expectedError:   "container is paused and no logs",
		},
		{
			scenario: "testable but not running",
			mockTarget: waitmock.MockStrategyTarget(func(t *waitmock.StrategyTarget) {
				t.On("State", isContext).
					Return(restartingState, nil)
			}),
			expectedSuccess: false,
			expectedError:   "health check failed: max retries exceeded",
		},
		{
			scenario: "testable but exec error",
			mockTarget: waitmock.MockStrategyTarget(func(t *waitmock.StrategyTarget) {
				t.On("State", isContext).
					Return(runningState, nil)

				t.On("Exec", isContext, []string{"test"}).
					Return(0, errors.New("exec error"))
			}),
			expectedSuccess: false,
			expectedError:   "exec error",
		},
		{
			scenario: "testable but return code not ok",
			mockTarget: waitmock.MockStrategyTarget(func(t *waitmock.StrategyTarget) {
				t.On("State", isContext).
					Return(runningState, nil)

				t.On("Exec", isContext, []string{"test"}).
					Return(1, nil)
			}),
			expectedSuccess: false,
			expectedError:   "health check failed: max retries exceeded",
		},
		{
			scenario: "testable and code ok",
			mockTarget: waitmock.MockStrategyTarget(func(t *waitmock.StrategyTarget) {
				t.On("State", isContext).
					Return(runningState, nil)

				t.On("Exec", isContext, []string{"test"}).
					Return(0, nil)
			}),
			expectedSuccess: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			s := wait.ForHealthCheckCmd("test").
				WithRetries(0).
				WithTestTimeout(time.Minute).
				WithStartPeriod(0)

			err := s.WaitUntilReady(context.Background(), tc.mockTarget(t))

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

var isContext = mock.MatchedBy(func(ctx interface{}) bool {
	_, is := ctx.(context.Context)

	return is
})

type ioReaderFunc func(b []byte) (n int, err error)

func (f ioReaderFunc) Read(b []byte) (n int, err error) {
	return f(b)
}

func errorReadCloser(err error) io.ReadCloser {
	return io.NopCloser(ioReaderFunc(func([]byte) (int, error) {
		return 0, err
	}))
}

func stringReadCloser(s string) io.ReadCloser {
	return io.NopCloser(ioReaderFunc(func(b []byte) (n int, err error) {
		if s == "" {
			return 0, io.EOF
		}

		l := len(s)
		c := cap(b)

		if c < l {
			l = c
		}

		copy(b, s[:l])
		s = s[l:]

		return l, nil
	}))
}
