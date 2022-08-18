package mock_test

import (
	"context"
	"errors"
	"io"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/assert"

	"go.nhat.io/testcontainers-go-extra/mock"
)

func TestContainer_GetContainerID(t *testing.T) {
	t.Parallel()

	c := mock.MockContainer(func(c *mock.Container) {
		c.On("GetContainerID").
			Return("id")
	})(t)

	actual := c.GetContainerID()
	expected := "id"

	assert.Equal(t, expected, actual)
}

func TestContainer_Endpoint(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mock           mock.ContainerMocker
		expectedResult string
		expectedError  error
	}{
		{
			scenario: "error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("Endpoint", context.Background(), "proto").
					Return("", errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "no error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("Endpoint", context.Background(), "proto").
					Return("endpoint", nil)
			}),
			expectedResult: "endpoint",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			result, err := tc.mock(t).Endpoint(context.Background(), "proto")

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestContainer_PortEndpoint(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mock           mock.ContainerMocker
		expectedResult string
		expectedError  error
	}{
		{
			scenario: "error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("PortEndpoint", context.Background(), nat.Port("8080"), "proto").
					Return("", errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "no error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("PortEndpoint", context.Background(), nat.Port("8080"), "proto").
					Return("endpoint", nil)
			}),
			expectedResult: "endpoint",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			result, err := tc.mock(t).PortEndpoint(context.Background(), "8080", "proto")

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestContainer_Host(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mock           mock.ContainerMocker
		expectedResult string
		expectedError  error
	}{
		{
			scenario: "error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("Host", context.Background()).
					Return("", errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "no error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("Host", context.Background()).
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

func TestContainer_MappedPort(t *testing.T) {
	t.Parallel()

	port := nat.Port("8080")

	testCases := []struct {
		scenario       string
		mock           mock.ContainerMocker
		expectedResult nat.Port
		expectedError  error
	}{
		{
			scenario: "string and error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("MappedPort", context.Background(), port).
					Return("", errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "string and no error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("MappedPort", context.Background(), port).
					Return("8080", nil)
			}),
			expectedResult: port,
		},
		{
			scenario: "port and error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("MappedPort", context.Background(), port).
					Return(nat.Port(""), errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "port and no error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("MappedPort", context.Background(), port).
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

func TestContainer_Ports(t *testing.T) {
	t.Parallel()

	port := nat.Port("8080")

	testCases := []struct {
		scenario       string
		mock           mock.ContainerMocker
		expectedResult nat.PortMap
		expectedError  error
	}{
		{
			scenario: "error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("Ports", context.Background()).
					Return(nil, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "no error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("Ports", context.Background()).
					Return(nat.PortMap{port: []nat.PortBinding{}}, nil)
			}),
			expectedResult: nat.PortMap{port: []nat.PortBinding{}},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			result, err := tc.mock(t).Ports(context.Background())

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestContainer_SessionID(t *testing.T) {
	t.Parallel()

	c := mock.MockContainer(func(c *mock.Container) {
		c.On("SessionID").
			Return("id")
	})(t)

	actual := c.SessionID()
	expected := "id"

	assert.Equal(t, expected, actual)
}

func TestContainer_Start(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mock          mock.ContainerMocker
		expectedError error
	}{
		{
			scenario: "error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("Start", context.Background()).
					Return(errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "no error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("Start", context.Background()).
					Return(nil)
			}),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			err := tc.mock(t).Start(context.Background())

			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestContainer_Stop(t *testing.T) {
	t.Parallel()

	duration := time.Second

	testCases := []struct {
		scenario      string
		mock          mock.ContainerMocker
		expectedError error
	}{
		{
			scenario: "error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("Stop", context.Background(), &duration).
					Return(errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "no error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("Stop", context.Background(), &duration).
					Return(nil)
			}),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			err := tc.mock(t).Stop(context.Background(), &duration)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestContainer_Terminate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mock          mock.ContainerMocker
		expectedError error
	}{
		{
			scenario: "error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("Terminate", context.Background()).
					Return(errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "no error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("Terminate", context.Background()).
					Return(nil)
			}),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			err := tc.mock(t).Terminate(context.Background())

			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestContainer_Logs(t *testing.T) {
	t.Parallel()

	rc := io.NopCloser(nil)

	testCases := []struct {
		scenario       string
		mock           mock.ContainerMocker
		expectedResult io.ReadCloser
		expectedError  error
	}{
		{
			scenario: "error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("Logs", context.Background()).
					Return(nil, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "no error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("Logs", context.Background()).
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

func TestContainer_FollowOutput(t *testing.T) {
	t.Parallel()

	c := mock.MockContainer(func(c *mock.Container) {
		c.On("FollowOutput", nil)
	})(t)

	c.FollowOutput(nil)
}

func TestContainer_StartLogProducer(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mock          mock.ContainerMocker
		expectedError error
	}{
		{
			scenario: "error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("StartLogProducer", context.Background()).
					Return(errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "no error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("StartLogProducer", context.Background()).
					Return(nil)
			}),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			err := tc.mock(t).StartLogProducer(context.Background())

			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestContainer_StopLogProducer(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mock          mock.ContainerMocker
		expectedError error
	}{
		{
			scenario: "error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("StopLogProducer").
					Return(errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "no error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("StopLogProducer").
					Return(nil)
			}),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			err := tc.mock(t).StopLogProducer()

			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestContainer_Name(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mock           mock.ContainerMocker
		expectedResult string
		expectedError  error
	}{
		{
			scenario: "error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("Name", context.Background()).
					Return("", errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "no error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("Name", context.Background()).
					Return("name", nil)
			}),
			expectedResult: "name",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			result, err := tc.mock(t).Name(context.Background())

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestContainer_State(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mock           mock.ContainerMocker
		expectedResult *types.ContainerState
		expectedError  error
	}{
		{
			scenario: "error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("State", context.Background()).
					Return(nil, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "no error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("State", context.Background()).
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

func TestContainer_Networks(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mock           mock.ContainerMocker
		expectedResult []string
		expectedError  error
	}{
		{
			scenario: "error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("Networks", context.Background()).
					Return(nil, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "no error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("Networks", context.Background()).
					Return([]string{"network"}, nil)
			}),
			expectedResult: []string{"network"},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			result, err := tc.mock(t).Networks(context.Background())

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestContainer_NetworkAliases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mock           mock.ContainerMocker
		expectedResult map[string][]string
		expectedError  error
	}{
		{
			scenario: "error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("NetworkAliases", context.Background()).
					Return(nil, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "no error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("NetworkAliases", context.Background()).
					Return(map[string][]string{"net1": {"network"}}, nil)
			}),
			expectedResult: map[string][]string{"net1": {"network"}},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			result, err := tc.mock(t).NetworkAliases(context.Background())

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestContainer_Exec(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mock           mock.ContainerMocker
		expectedResult int
		expectedError  error
	}{
		{
			scenario: "error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("Exec", context.Background(), []string{"test"}).
					Return(1, errors.New("error"))
			}),
			expectedResult: 1,
			expectedError:  errors.New("error"),
		},
		{
			scenario: "no error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("Exec", context.Background(), []string{"test"}).
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

func TestContainer_ContainerIP(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		mock           mock.ContainerMocker
		expectedResult string
		expectedError  error
	}{
		{
			scenario: "error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("ContainerIP", context.Background()).
					Return("", errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "no error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("ContainerIP", context.Background()).
					Return("ip", nil)
			}),
			expectedResult: "ip",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			result, err := tc.mock(t).ContainerIP(context.Background())

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestContainer_CopyToContainer(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mock          mock.ContainerMocker
		expectedError error
	}{
		{
			scenario: "error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("CopyToContainer", context.Background(), []byte(`hello world`), "/tmp/test.csv", int64(0)).
					Return(errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "no error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("CopyToContainer", context.Background(), []byte(`hello world`), "/tmp/test.csv", int64(0)).
					Return(nil)
			}),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			err := tc.mock(t).CopyToContainer(context.Background(), []byte(`hello world`), "/tmp/test.csv", 0)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestContainer_CopyFileToContainer(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		mock          mock.ContainerMocker
		expectedError error
	}{
		{
			scenario: "error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("CopyFileToContainer", context.Background(), "test.csv", "/tmp/test.csv", int64(0)).
					Return(errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "no error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("CopyFileToContainer", context.Background(), "test.csv", "/tmp/test.csv", int64(0)).
					Return(nil)
			}),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			err := tc.mock(t).CopyFileToContainer(context.Background(), "test.csv", "/tmp/test.csv", 0)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestContainer_CopyFileFromContainer(t *testing.T) {
	t.Parallel()

	rc := io.NopCloser(nil)

	testCases := []struct {
		scenario       string
		mock           mock.ContainerMocker
		expectedResult io.ReadCloser
		expectedError  error
	}{
		{
			scenario: "error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("CopyFileFromContainer", context.Background(), "/tmp/test.csv").
					Return(nil, errors.New("error"))
			}),
			expectedError: errors.New("error"),
		},
		{
			scenario: "no error",
			mock: mock.MockContainer(func(c *mock.Container) {
				c.On("CopyFileFromContainer", context.Background(), "/tmp/test.csv").
					Return(rc, nil)
			}),
			expectedResult: rc,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			result, err := tc.mock(t).CopyFileFromContainer(context.Background(), "/tmp/test.csv")

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
