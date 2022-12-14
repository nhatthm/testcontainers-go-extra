// Code generated by mockery v2.14.0. DO NOT EDIT.

package mock

import (
	context "context"
	io "io"

	exec "github.com/testcontainers/testcontainers-go/exec"

	mock "github.com/stretchr/testify/mock"

	nat "github.com/docker/go-connections/nat"

	testcontainers "github.com/testcontainers/testcontainers-go"

	time "time"

	types "github.com/docker/docker/api/types"
)

// Container is an autogenerated mock type for the Container type
type Container struct {
	mock.Mock
}

// ContainerIP provides a mock function with given fields: _a0
func (_m *Container) ContainerIP(_a0 context.Context) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerIPs provides a mock function with given fields: _a0
func (_m *Container) ContainerIPs(_a0 context.Context) ([]string, error) {
	ret := _m.Called(_a0)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CopyDirToContainer provides a mock function with given fields: ctx, hostDirPath, containerParentPath, fileMode
func (_m *Container) CopyDirToContainer(ctx context.Context, hostDirPath string, containerParentPath string, fileMode int64) error {
	ret := _m.Called(ctx, hostDirPath, containerParentPath, fileMode)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) error); ok {
		r0 = rf(ctx, hostDirPath, containerParentPath, fileMode)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CopyFileFromContainer provides a mock function with given fields: ctx, filePath
func (_m *Container) CopyFileFromContainer(ctx context.Context, filePath string) (io.ReadCloser, error) {
	ret := _m.Called(ctx, filePath)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, string) io.ReadCloser); ok {
		r0 = rf(ctx, filePath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CopyFileToContainer provides a mock function with given fields: ctx, hostFilePath, containerFilePath, fileMode
func (_m *Container) CopyFileToContainer(ctx context.Context, hostFilePath string, containerFilePath string, fileMode int64) error {
	ret := _m.Called(ctx, hostFilePath, containerFilePath, fileMode)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) error); ok {
		r0 = rf(ctx, hostFilePath, containerFilePath, fileMode)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CopyToContainer provides a mock function with given fields: ctx, fileContent, containerFilePath, fileMode
func (_m *Container) CopyToContainer(ctx context.Context, fileContent []byte, containerFilePath string, fileMode int64) error {
	ret := _m.Called(ctx, fileContent, containerFilePath, fileMode)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, string, int64) error); ok {
		r0 = rf(ctx, fileContent, containerFilePath, fileMode)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Endpoint provides a mock function with given fields: _a0, _a1
func (_m *Container) Endpoint(_a0 context.Context, _a1 string) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Exec provides a mock function with given fields: ctx, cmd, options
func (_m *Container) Exec(ctx context.Context, cmd []string, options ...exec.ProcessOption) (int, io.Reader, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, cmd)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, []string, ...exec.ProcessOption) int); ok {
		r0 = rf(ctx, cmd, options...)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 io.Reader
	if rf, ok := ret.Get(1).(func(context.Context, []string, ...exec.ProcessOption) io.Reader); ok {
		r1 = rf(ctx, cmd, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(io.Reader)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, []string, ...exec.ProcessOption) error); ok {
		r2 = rf(ctx, cmd, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FollowOutput provides a mock function with given fields: _a0
func (_m *Container) FollowOutput(_a0 testcontainers.LogConsumer) {
	_m.Called(_a0)
}

// GetContainerID provides a mock function with given fields:
func (_m *Container) GetContainerID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Host provides a mock function with given fields: _a0
func (_m *Container) Host(_a0 context.Context) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsRunning provides a mock function with given fields:
func (_m *Container) IsRunning() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Logs provides a mock function with given fields: _a0
func (_m *Container) Logs(_a0 context.Context) (io.ReadCloser, error) {
	ret := _m.Called(_a0)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context) io.ReadCloser); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MappedPort provides a mock function with given fields: _a0, _a1
func (_m *Container) MappedPort(_a0 context.Context, _a1 nat.Port) (nat.Port, error) {
	ret := _m.Called(_a0, _a1)

	var r0 nat.Port
	if rf, ok := ret.Get(0).(func(context.Context, nat.Port) nat.Port); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(nat.Port)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, nat.Port) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Name provides a mock function with given fields: _a0
func (_m *Container) Name(_a0 context.Context) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NetworkAliases provides a mock function with given fields: _a0
func (_m *Container) NetworkAliases(_a0 context.Context) (map[string][]string, error) {
	ret := _m.Called(_a0)

	var r0 map[string][]string
	if rf, ok := ret.Get(0).(func(context.Context) map[string][]string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Networks provides a mock function with given fields: _a0
func (_m *Container) Networks(_a0 context.Context) ([]string, error) {
	ret := _m.Called(_a0)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PortEndpoint provides a mock function with given fields: _a0, _a1, _a2
func (_m *Container) PortEndpoint(_a0 context.Context, _a1 nat.Port, _a2 string) (string, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, nat.Port, string) string); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, nat.Port, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Ports provides a mock function with given fields: _a0
func (_m *Container) Ports(_a0 context.Context) (nat.PortMap, error) {
	ret := _m.Called(_a0)

	var r0 nat.PortMap
	if rf, ok := ret.Get(0).(func(context.Context) nat.PortMap); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(nat.PortMap)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SessionID provides a mock function with given fields:
func (_m *Container) SessionID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *Container) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartLogProducer provides a mock function with given fields: _a0
func (_m *Container) StartLogProducer(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// State provides a mock function with given fields: _a0
func (_m *Container) State(_a0 context.Context) (*types.ContainerState, error) {
	ret := _m.Called(_a0)

	var r0 *types.ContainerState
	if rf, ok := ret.Get(0).(func(context.Context) *types.ContainerState); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ContainerState)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Stop provides a mock function with given fields: _a0, _a1
func (_m *Container) Stop(_a0 context.Context, _a1 *time.Duration) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *time.Duration) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StopLogProducer provides a mock function with given fields:
func (_m *Container) StopLogProducer() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Terminate provides a mock function with given fields: _a0
func (_m *Container) Terminate(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewContainer interface {
	mock.TestingT
	Cleanup(func())
}

// NewContainer creates a new instance of Container. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewContainer(t mockConstructorTestingTNewContainer) *Container {
	mock := &Container{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
