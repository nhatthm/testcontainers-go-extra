package mock

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/testcontainers/testcontainers-go"
)

// ContainerMocker is Container mocker.
type ContainerMocker func(tb testing.TB) *Container

// NoMockContainer is no mock Container.
var NoMockContainer = MockContainer()

var _ testcontainers.Container = (*Container)(nil)

// Container is a testcontainers.Container.
type Container struct {
	mock.Mock
}

// GetContainerID satisfies testcontainers.Container interface.
func (c *Container) GetContainerID() string {
	return c.Called().String(0)
}

// Endpoint satisfies testcontainers.Container interface.
func (c *Container) Endpoint(ctx context.Context, proto string) (string, error) {
	result := c.Called(ctx, proto)

	return result.String(0), result.Error(1)
}

// PortEndpoint satisfies testcontainers.Container interface.
func (c *Container) PortEndpoint(ctx context.Context, port nat.Port, proto string) (string, error) {
	result := c.Called(ctx, port, proto)

	return result.String(0), result.Error(1)
}

// Host satisfies testcontainers.Container interface.
func (c *Container) Host(ctx context.Context) (string, error) {
	result := c.Called(ctx)

	return result.String(0), result.Error(1)
}

// MappedPort satisfies testcontainers.Container interface.
func (c *Container) MappedPort(ctx context.Context, port nat.Port) (nat.Port, error) {
	result := c.Called(ctx, port)

	p := result.Get(0)
	err := result.Error(1)

	if s, ok := p.(string); ok {
		return nat.Port(s), err
	}

	return p.(nat.Port), err
}

// Ports satisfies testcontainers.Container interface.
func (c *Container) Ports(ctx context.Context) (nat.PortMap, error) {
	result := c.Called(ctx)

	m := result.Get(0)
	err := result.Error(1)

	if m == nil {
		return nil, err
	}

	return m.(nat.PortMap), err
}

// SessionID satisfies testcontainers.Container interface.
func (c *Container) SessionID() string {
	return c.Called().String(0)
}

// Start satisfies testcontainers.Container interface.
func (c *Container) Start(ctx context.Context) error {
	return c.Called(ctx).Error(0)
}

// Stop satisfies testcontainers.Container interface.
func (c *Container) Stop(ctx context.Context, duration *time.Duration) error {
	return c.Called(ctx, duration).Error(0)
}

// Terminate satisfies testcontainers.Container interface.
func (c *Container) Terminate(ctx context.Context) error {
	return c.Called(ctx).Error(0)
}

// Logs satisfies testcontainers.Container interface.
func (c *Container) Logs(ctx context.Context) (io.ReadCloser, error) {
	result := c.Called(ctx)

	rc := result.Get(0)
	err := result.Error(1)

	if rc == nil {
		return nil, err
	}

	return rc.(io.ReadCloser), err
}

// FollowOutput satisfies testcontainers.Container interface.
func (c *Container) FollowOutput(consumer testcontainers.LogConsumer) {
	c.Called(consumer)
}

// StartLogProducer satisfies testcontainers.Container interface.
func (c *Container) StartLogProducer(ctx context.Context) error {
	return c.Called(ctx).Error(0)
}

// StopLogProducer satisfies testcontainers.Container interface.
func (c *Container) StopLogProducer() error {
	return c.Called().Error(0)
}

// Name satisfies testcontainers.Container interface.
func (c *Container) Name(ctx context.Context) (string, error) {
	result := c.Called(ctx)

	return result.String(0), result.Error(1)
}

// State satisfies testcontainers.Container interface.
func (c *Container) State(ctx context.Context) (*types.ContainerState, error) {
	result := c.Called(ctx)

	state := result.Get(0)
	err := result.Error(1)

	if state == nil {
		return nil, err
	}

	return state.(*types.ContainerState), err
}

// Networks satisfies testcontainers.Container interface.
func (c *Container) Networks(ctx context.Context) ([]string, error) {
	result := c.Called(ctx)

	networks := result.Get(0)
	err := result.Error(1)

	if networks == nil {
		return nil, err
	}

	return networks.([]string), err
}

// NetworkAliases satisfies testcontainers.Container interface.
func (c *Container) NetworkAliases(ctx context.Context) (map[string][]string, error) {
	result := c.Called(ctx)

	networks := result.Get(0)
	err := result.Error(1)

	if networks == nil {
		return nil, err
	}

	return networks.(map[string][]string), err
}

// Exec satisfies testcontainers.Container interface.
func (c *Container) Exec(ctx context.Context, cmd []string) (int, io.Reader, error) {
	result := c.Called(ctx, cmd)

	r1, r2, r3 := result.Int(0), result.Get(1), result.Error(2)

	if r2 == nil {
		return r1, nil, r3
	}

	return r1, r2.(io.Reader), r3
}

// ContainerIP satisfies testcontainers.Container interface.
func (c *Container) ContainerIP(ctx context.Context) (string, error) {
	result := c.Called(ctx)

	return result.String(0), result.Error(1)
}

// CopyToContainer satisfies testcontainers.Container interface.
func (c *Container) CopyToContainer(ctx context.Context, fileContent []byte, containerFilePath string, fileMode int64) error {
	return c.Called(ctx, fileContent, containerFilePath, fileMode).Error(0)
}

// CopyFileToContainer satisfies testcontainers.Container interface.
func (c *Container) CopyFileToContainer(ctx context.Context, hostFilePath string, containerFilePath string, fileMode int64) error {
	return c.Called(ctx, hostFilePath, containerFilePath, fileMode).Error(0)
}

// CopyDirToContainer satisfies testcontainers.Container interface.
func (c *Container) CopyDirToContainer(ctx context.Context, hostDirPath string, containerParentPath string, fileMode int64) error {
	return c.Called(ctx, hostDirPath, containerParentPath, fileMode).Error(0)
}

// CopyFileFromContainer satisfies testcontainers.Container interface.
func (c *Container) CopyFileFromContainer(ctx context.Context, filePath string) (io.ReadCloser, error) {
	result := c.Called(ctx, filePath)

	rc := result.Get(0)
	err := result.Error(1)

	if rc == nil {
		return nil, err
	}

	return rc.(io.ReadCloser), err
}

// IsRunning satisfies testcontainers.Container interface.
func (c *Container) IsRunning() bool {
	return c.Called().Bool(0)
}

// mockContainer mocks testcontainers.Container interface.
func mockContainer(mocks ...func(c *Container)) *Container {
	c := &Container{}

	for _, m := range mocks {
		m(c)
	}

	return c
}

// MockContainer creates Container mock with cleanup to ensure all the expectations are met.
// nolint: revive
func MockContainer(mocks ...func(c *Container)) ContainerMocker {
	return func(tb testing.TB) *Container {
		tb.Helper()

		c := mockContainer(mocks...)

		tb.Cleanup(func() {
			assert.True(tb, c.Mock.AssertExpectations(tb))
		})

		return c
	}
}
