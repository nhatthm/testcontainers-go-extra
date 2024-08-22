package testcontainers_test

import (
	"context"
	crand "crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nhat.io/testcontainers-extra"
	"go.nhat.io/testcontainers-extra/wait"
)

func TestStartGenericContainer_GenericContainerError(t *testing.T) {
	t.Parallel()

	c, err := testcontainers.StartGenericContainer(
		context.Background(),
		testcontainers.ContainerRequest{},
		testcontainers.WithProviderType(-1),
	)

	expected := `get provider: unknown provider`

	assert.Nil(t, c)
	assert.EqualError(t, err, expected)
}

func TestStartGenericContainer_Success(t *testing.T) {
	t.Parallel()

	c, err := testcontainers.StartGenericContainer(context.Background(), testcontainers.ContainerRequest{
		Name:  randomString(8),
		Image: "alpine",
	})

	require.NotNil(t, c)
	assert.NoError(t, err)

	defer c.Terminate(context.Background()) // nolint: errcheck
}

func TestStartGenericContainer_SuccessButCallbackError(t *testing.T) {
	t.Parallel()

	c, err := testcontainers.StartGenericContainer(context.Background(), testcontainers.ContainerRequest{
		Name:  randomString(8),
		Image: "alpine",
	}, testcontainers.WithCallback(func(context.Context, testcontainers.Container, testcontainers.ContainerRequest) error {
		return errors.New("callback error")
	}))

	require.NotNil(t, c)

	defer c.Terminate(context.Background()) // nolint: errcheck

	expected := errors.New("callback error")

	assert.Equal(t, expected, err)
}

func TestStartGenericContainer_NamePrefix(t *testing.T) {
	t.Parallel()

	prefix := randomString(8)

	c, err := testcontainers.StartGenericContainer(context.Background(), testcontainers.ContainerRequest{
		Name:  "test",
		Image: "alpine",
	}, testcontainers.WithNamePrefix(prefix))

	require.NotNil(t, c)
	assert.NoError(t, err)

	defer c.Terminate(context.Background()) // nolint: errcheck

	actual, err := c.Name(context.Background())
	expected := fmt.Sprintf("/%s_test", prefix)

	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestStartGenericContainer_NameSuffix(t *testing.T) {
	t.Parallel()

	suffix := randomString(8)

	c, err := testcontainers.StartGenericContainer(context.Background(), testcontainers.ContainerRequest{
		Name:  "test",
		Image: "alpine",
	}, testcontainers.WithNameSuffix(suffix))

	require.NotNil(t, c)
	assert.NoError(t, err)

	defer c.Terminate(context.Background()) // nolint: errcheck

	actual, err := c.Name(context.Background())
	expected := fmt.Sprintf("/test_%s", suffix)

	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestStartGenericContainer_NamePrefixAndSuffix(t *testing.T) {
	t.Parallel()

	prefix := randomString(8)
	suffix := randomString(8)

	c, err := testcontainers.StartGenericContainer(context.Background(), testcontainers.ContainerRequest{
		Name:  "test",
		Image: "alpine",
	}, testcontainers.WithNamePrefix(prefix), testcontainers.WithNameSuffix(suffix))

	require.NotNil(t, c)
	assert.NoError(t, err)

	defer c.Terminate(context.Background()) // nolint: errcheck

	actual, err := c.Name(context.Background())
	expected := fmt.Sprintf("/%s_test_%s", prefix, suffix)

	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestStartGenericContainer_WithImageName(t *testing.T) {
	t.Parallel()

	c, err := testcontainers.StartGenericContainer(context.Background(), testcontainers.ContainerRequest{
		Name:  randomString(8),
		Image: "alpine-unknown",
	}, testcontainers.WithImageName("alpine"))

	require.NotNil(t, c)
	assert.NoError(t, err)

	defer c.Terminate(context.Background()) // nolint: errcheck
}

func TestStartGenericContainer_WithImageTag(t *testing.T) {
	t.Parallel()

	c, err := testcontainers.StartGenericContainer(context.Background(), testcontainers.ContainerRequest{
		Name:  randomString(8),
		Image: "alpine:unknown",
	}, testcontainers.WithImageTag("latest"))

	require.NotNil(t, c)
	assert.NoError(t, err)

	defer c.Terminate(context.Background()) // nolint: errcheck
}

func TestStartGenericContainers_Success(t *testing.T) {
	t.Parallel()

	containers, err := testcontainers.StartGenericContainers(context.Background(),
		testcontainers.StartGenericContainerRequest{
			Request: testcontainers.ContainerRequest{
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
			},
			Options: []testcontainers.GenericContainerOption{
				testcontainers.WithNameSuffix(randomString(8)),
			},
		},
		testcontainers.StartGenericContainerRequest{
			Request: testcontainers.ContainerRequest{
				Name:  randomString(8),
				Image: "alpine",
			},
		},
	)

	assert.NoError(t, err)
	require.Len(t, containers, 2)
}

func TestStartGenericContainers_Error(t *testing.T) {
	t.Parallel()

	containers, err := testcontainers.StartGenericContainers(context.Background(),
		testcontainers.StartGenericContainerRequest{
			Request: testcontainers.ContainerRequest{
				Name: randomString(8),
			},
			Options: []testcontainers.GenericContainerOption{
				testcontainers.WithProviderType(-1),
			},
		},
		testcontainers.StartGenericContainerRequest{
			Request: testcontainers.ContainerRequest{
				Name: randomString(8),
			},
			Options: []testcontainers.GenericContainerOption{
				testcontainers.WithProviderType(-1),
			},
		},
		testcontainers.StartGenericContainerRequest{
			Request: testcontainers.ContainerRequest{
				Name:  randomString(8),
				Image: "alpine",
			},
		},
	)

	t.Logf("logs:\n%s", err.Error())

	require.NotEmpty(t, containers)
}

func TestStopGenericContainers_Success(t *testing.T) {
	t.Parallel()

	containers, err := testcontainers.StartGenericContainers(context.Background(),
		testcontainers.StartGenericContainerRequest{
			Request: testcontainers.ContainerRequest{
				Name:  randomString(8),
				Image: "alpine",
			},
		},
	)

	require.NoError(t, err)
	require.NotEmpty(t, containers)

	err = testcontainers.StopGenericContainers(context.Background(), containers...)

	assert.NoError(t, err)
}

func TestStopGenericContainers_Error(t *testing.T) {
	t.Parallel()

	containers, err := testcontainers.StartGenericContainers(context.Background(),
		testcontainers.StartGenericContainerRequest{
			Request: testcontainers.ContainerRequest{
				Name:  randomString(8),
				Image: "alpine",
			},
		},
	)

	require.NotEmpty(t, containers)
	require.NoError(t, err)

	_ = containers[0].Terminate(context.Background()) // nolint: errcheck

	err = testcontainers.StopGenericContainers(context.Background(), containers...)

	require.Error(t, err)

	t.Logf("logs:\n%s", err.Error())
}

// nolint: unparam
func randomString(length int) string {
	var rngSeed int64

	_ = binary.Read(crand.Reader, binary.LittleEndian, &rngSeed) // nolint: errcheck
	r := rand.New(rand.NewSource(rngSeed))                       // nolint: gosec

	result := make([]byte, length/2)

	_, _ = r.Read(result)

	return hex.EncodeToString(result)
}
