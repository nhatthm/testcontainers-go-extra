package testcontainers

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/docker/go-connections/nat"
)

var setEnv = os.Setenv

// PopulateHostPortEnv sets the hostname and public port for each exposed port.
var PopulateHostPortEnv = ContainerCallback(func(ctx context.Context, c Container, r ContainerRequest) error {
	ports, err := c.Ports(ctx)
	if err != nil {
		return fmt.Errorf("could not get container %q ports: %w", r.Name, err)
	}

	if len(ports) == 0 {
		return nil
	}

	ip, err := c.Host(ctx)
	if err != nil {
		return fmt.Errorf("could not get container %q ip: %w", r.Name, err)
	}

	for p, bindings := range ports {
		for _, b := range bindings {
			if err := setEnvVar(r.Name, p, "HOST", ip); err != nil {
				return err
			}

			if err := setEnvVar(r.Name, p, "PORT", b.HostPort); err != nil {
				return err
			}
		}
	}

	return nil
})

var alphaNum = regexp.MustCompile("[^a-zA-Z0-9]+")

func envVarName(parts ...string) string {
	return strings.ToUpper(alphaNum.ReplaceAllString(strings.Join(parts, "_"), "_"))
}

func setEnvVar(containerName string, containerPort nat.Port, suffix, value string) error {
	envVar := envVarName(containerName, containerPort.Port(), suffix)

	if err := setEnv(envVar, value); err != nil {
		return fmt.Errorf("could not set env var %q: %w", envVar, err)
	}

	return nil
}
