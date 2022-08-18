package wait_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/testcontainers-extra/mock/wait"
)

func TestStrategy_WaitUntilReady(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario string
		mock     wait.StrategyMocker
		expected error
	}{
		{
			scenario: "error",
			mock: wait.MockStrategy(func(s *wait.Strategy) {
				s.On("WaitUntilReady", context.Background(), nil).
					Return(errors.New("error"))
			}),
			expected: errors.New("error"),
		},
		{
			scenario: "no error",
			mock: wait.MockStrategy(func(s *wait.Strategy) {
				s.On("WaitUntilReady", context.Background(), nil).
					Return(nil)
			}),
			expected: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual := tc.mock(t).WaitUntilReady(context.Background(), nil)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
