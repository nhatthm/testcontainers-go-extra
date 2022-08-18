//go:build go1.17
// +build go1.17

package testcontainers

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/assert"

	"go.nhat.io/testcontainers-extra/mock"
)

// nolint: paralleltest
func TestPopulateHostPortEnv(t *testing.T) {
	// Do not run this test in parallel because we can't test the os.Setenv.
	ports := nat.PortMap{
		"8080/tcp": {
			{HostIP: "0.0.0.0", HostPort: "8080"},
		},
	}

	request := ContainerRequest{Name: "test"}

	testCases := []struct {
		scenario      string
		mockContainer mock.ContainerMocker
		setEnv        func(key, value string) error
		request       ContainerRequest
		expectedHost  string
		expectedPort  string
		expectedError string
	}{
		{
			scenario: "could not get ports",
			mockContainer: mock.MockContainer(func(c *mock.Container) {
				c.On("Ports", context.Background()).
					Return(nil, errors.New("get ports error"))
			}),
			request:       request,
			expectedError: `could not get container "test" ports: get ports error`,
		},
		{
			scenario: "no ports",
			mockContainer: mock.MockContainer(func(c *mock.Container) {
				c.On("Ports", context.Background()).
					Return(nil, nil)
			}),
			request: request,
		},
		{
			scenario: "could not get host",
			mockContainer: mock.MockContainer(func(c *mock.Container) {
				c.On("Ports", context.Background()).
					Return(ports, nil)

				c.On("Host", context.Background()).
					Return("", errors.New("get host error"))
			}),
			request:       request,
			expectedError: `could not get container "test" ip: get host error`,
		},
		{
			scenario: "could not set host",
			mockContainer: mock.MockContainer(func(c *mock.Container) {
				c.On("Ports", context.Background()).
					Return(ports, nil)

				c.On("Host", context.Background()).
					Return("localhost", nil)
			}),
			setEnv: func(string, string) error {
				return errors.New("set error")
			},
			request:       request,
			expectedError: `could not set env var "TEST_8080_HOST": set error`,
		},
		{
			scenario: "could not set port",
			mockContainer: mock.MockContainer(func(c *mock.Container) {
				c.On("Ports", context.Background()).
					Return(ports, nil)

				c.On("Host", context.Background()).
					Return("localhost", nil)
			}),
			setEnv: func(k string, _ string) error {
				if strings.Contains(k, "HOST") {
					return nil
				}

				return errors.New("set error")
			},
			request:       request,
			expectedError: `could not set env var "TEST_8080_PORT": set error`,
		},
		{
			scenario: "success",
			mockContainer: mock.MockContainer(func(c *mock.Container) {
				c.On("Ports", context.Background()).
					Return(ports, nil)

				c.On("Host", context.Background()).
					Return("localhost", nil)
			}),
			request:      request,
			expectedHost: "localhost",
			expectedPort: "8080",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			if tc.setEnv == nil {
				setEnvTest(t)
			} else {
				setEnv = tc.setEnv
			}

			err := PopulateHostPortEnv(context.Background(), tc.mockContainer(t), tc.request)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			host := os.Getenv("TEST_8080_HOST")
			port := os.Getenv("TEST_8080_PORT")

			assert.Equal(t, tc.expectedHost, host)
			assert.Equal(t, tc.expectedPort, port)
		})
	}
}

func setEnvTest(t *testing.T) {
	t.Helper()

	setEnv = func(key, value string) (err error) {
		defer func() {
			if r := recover(); r != nil {
				if re, ok := r.(error); ok {
					err = re
				} else {
					err = fmt.Errorf("%+v", r)
				}
			}
		}()

		t.Setenv(key, value)

		return nil
	}

	t.Cleanup(func() {
		setEnv = os.Setenv
	})
}
