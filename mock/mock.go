package mock

import (
	"testing"

	"github.com/testcontainers/testcontainers-go"
)

//go:generate bash -c "mockery --name Container --dir=\"$(go env GOMODCACHE)/github.com/testcontainers/testcontainers-go@$(go list -m -f '{{ .Version }}' github.com/testcontainers/testcontainers-go)\" --output . --outpkg mock --filename container.go"

// ContainerMocker is Container mocker.
type ContainerMocker func(tb testing.TB) *Container

// NopContainer is no mock Container.
var NopContainer = MockContainer()

var _ testcontainers.Container = (*Container)(nil)

// MockContainer creates Container mock with cleanup to ensure all the expectations are met.
// nolint: revive
func MockContainer(mocks ...func(c *Container)) ContainerMocker {
	return func(tb testing.TB) *Container {
		tb.Helper()

		c := NewContainer(tb)

		for _, m := range mocks {
			m(c)
		}

		return c
	}
}
