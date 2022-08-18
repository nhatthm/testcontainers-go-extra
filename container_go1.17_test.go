//go:build go1.17
// +build go1.17

package testcontainers_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nhat.io/testcontainers-extra"
	"go.nhat.io/testcontainers-extra/wait"
)

// nolint: paralleltest
func TestStartGenericContainer_PopulateHostPortEnv(t *testing.T) {
	t.Setenv("POSTGRES_5432_HOST", "")
	t.Setenv("POSTGRES_5432_PORT", "")

	c, err := testcontainers.StartGenericContainer(context.Background(), testcontainers.ContainerRequest{
		Name:         "postgres",
		Image:        "postgres:12-alpine",
		ExposedPorts: []string{":5432"},
		Env: map[string]string{
			"LC_ALL":            "C.UTF-8",
			"POSTGRES_DB":       "test",
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "test",
		},
		WaitingFor: wait.ForHealthCheckCmd("pg_isready").
			WithRetries(3).
			WithStartPeriod(time.Second).
			WithTestTimeout(5 * time.Second).
			WithTestInterval(10 * time.Second),
	}, testcontainers.PopulateHostPortEnv, testcontainers.WithNameSuffix(randomString(8)))

	require.NotNil(t, c)
	assert.NoError(t, err)

	defer c.Terminate(context.Background()) // nolint: errcheck

	host := os.Getenv("POSTGRES_5432_HOST")
	port := os.Getenv("POSTGRES_5432_PORT")

	assert.NotEmpty(t, host)
	assert.NotEmpty(t, port)
}
