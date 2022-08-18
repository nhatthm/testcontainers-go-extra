package wait

import (
	"context"
	"fmt"
	"io"

	"github.com/testcontainers/testcontainers-go/wait"

	"go.nhat.io/testcontainers-extra"
)

func healthCheckTestCmd(cmd []string) HealthCheckTestFunc {
	return func(ctx context.Context, target wait.StrategyTarget) (success bool, err error) {
		state, err := target.State(ctx)
		if err != nil {
			return false, err
		}

		if !isCmdTestable(state.Status) {
			logs, err := target.Logs(ctx)
			if err != nil {
				return false, fmt.Errorf("container is %s and unable to get logs: %w", state.Status, err)
			}

			if logs != nil {
				out, err := io.ReadAll(logs)
				if err != nil {
					return false, fmt.Errorf("container is %s and unable to read logs: %w", state.Status, err)
				}

				return false, fmt.Errorf("container is %s, logs:\n%s", state.Status, string(out))
			}

			return false, fmt.Errorf("container is %s and no logs", state.Status)
		}

		if !state.Running {
			return false, nil
		}

		code, err := target.Exec(ctx, cmd)
		if err != nil {
			return false, err
		}

		return code == 0, nil
	}
}

// ForHealthCheckCmd checks by running a command in the container.
func ForHealthCheckCmd(cmd string, args ...string) *HealthCheckStrategy {
	test := make([]string, 0, len(args)+1)
	test = append(test, cmd)
	test = append(test, args...)

	return ForHealthCheck(healthCheckTestCmd(test))
}

func isCmdTestable(status string) bool {
	return testcontainers.ContainerStatusCreated.Equal(status) ||
		testcontainers.ContainerStatusRunning.Equal(status) ||
		testcontainers.ContainerStatusRestarting.Equal(status)
}
