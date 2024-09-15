# Testcontainers-Go Booster 

[![GitHub Releases](https://img.shields.io/github/v/release/nhatthm/testcontainers-go-extra)](https://github.com/nhatthm/testcontainers-go-extra/releases/latest)
[![Build Status](https://github.com/nhatthm/testcontainers-go-extra/actions/workflows/test.yaml/badge.svg)](https://github.com/nhatthm/testcontainers-go-extra/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/nhatthm/testcontainers-go-extra/branch/master/graph/badge.svg?token=eTdAgDE2vR)](https://codecov.io/gh/nhatthm/testcontainers-go-extra)
[![Go Report Card](https://goreportcard.com/badge/go.nhat.io/testcontainers-extra)](https://goreportcard.com/report/go.nhat.io/testcontainers-extra)
[![GoDevDoc](https://img.shields.io/badge/dev-doc-00ADD8?logo=go)](https://pkg.go.dev/go.nhat.io/testcontainers-extra)
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/donate/?hosted_button_id=PJZSGJN57TDJY)

Boost `testcontainers/testcontainers-go` with some jet fuel! 🚀

## Prerequisites

- `Go >= 1.22`

## Install

```bash
go get go.nhat.io/testcontainers-extra
```

## Callbacks

After successfully starting a container, you may want to do some extra operations to make it ready for your purposes.
Then you could use the callbacks.

For example:

```go
package example

import (
	"context"
	"time"

	"go.nhat.io/testcontainers-extra"
	"go.nhat.io/testcontainers-extra/wait"
)

func startPostgres(dbName, dbUser, dbPassword string) error {
	_, err := testcontainers.StartGenericContainer(context.Background(), testcontainers.ContainerRequest{
		Name:         "postgres",
		Image:        "postgres:12-alpine",
		ExposedPorts: []string{":5432"},
		Env: map[string]string{
			"LC_ALL":            "C.UTF-8",
			"POSTGRES_DB":       dbName,
			"POSTGRES_USER":     dbUser,
			"POSTGRES_PASSWORD": dbPassword,
		},
		WaitingFor: wait.ForHealthCheckCmd("pg_isready").
			WithRetries(3).
			WithStartPeriod(30 * time.Second).
			WithTestTimeout(5 * time.Second).
			WithTestInterval(10 * time.Second),
	}, testcontainers.WithCallback(func(ctx context.Context, c testcontainers.Container, r testcontainers.ContainerRequest) error {
		// Do your stuff here, for example, migration.

		return nil
	}))

	return err
}
```

### Populating Host and Ports Envs

`testcontainers.PopulateHostPortEnv` is a callback that set the environment variables for the exposed ports.

For example:

```go
package example

import (
	"context"
	"time"

	"go.nhat.io/testcontainers-extra"
	"go.nhat.io/testcontainers-extra/wait"
)

func startPostgres(dbName, dbUser, dbPassword string) error {
	_, err := testcontainers.StartGenericContainer(context.Background(), testcontainers.ContainerRequest{
		Name:         "postgres",
		Image:        "postgres:12-alpine",
		ExposedPorts: []string{":5432"},
		Env: map[string]string{
			"LC_ALL":            "C.UTF-8",
			"POSTGRES_DB":       dbName,
			"POSTGRES_USER":     dbUser,
			"POSTGRES_PASSWORD": dbPassword,
		},
		WaitingFor: wait.ForHealthCheckCmd("pg_isready").
			WithRetries(3).
			WithStartPeriod(30 * time.Second).
			WithTestTimeout(5 * time.Second).
			WithTestInterval(10 * time.Second),
	}, testcontainers.PopulateHostPortEnv)

	return err
}
```

After calling `startPostgres()`, there will be 2 variables `POSTGRES_5432_HOST` and `POSTGRES_5432_PORT`. The `POSTGRES`
is from the container name in the request, `5432` is the exposed port. The values are
- `POSTGRES_5432_HOST`: the hostname of the docker daemon where the container port is exposed. 
- `POSTGRES_5432_PORT`: the port that mapped to the exposed container port.

## Wait Strategies

### Health Check

The health check provides the same behavior as `docker-compose` with the configuration:

- `Start Period`: Retry is only counted when time passes the start period. This is helpful for some containers that need
  time to get ready. The default value is `0`.
- `Test Timeout`: Timeout for executing the test.
- `Test Internal`: If the container is unhealthy, the health check will wait for an amount of time before testing again.
- `Retries`: The number of retries to test the container after start period ends.

![hccmd](https://user-images.githubusercontent.com/1154587/151780048-558853c4-5395-4ae2-939c-a32d2306cf9a.png)

For example:

```go
package example

import (
	"context"
	"time"

	"go.nhat.io/testcontainers-extra"
	"go.nhat.io/testcontainers-extra/wait"
)

func startPostgres(dbName, dbUser, dbPassword string) error {
	_, err := testcontainers.StartGenericContainer(context.Background(), testcontainers.ContainerRequest{
		Name:         "postgres",
		Image:        "postgres:12-alpine",
		ExposedPorts: []string{":5432"},
		Env: map[string]string{
			"LC_ALL":            "C.UTF-8",
			"POSTGRES_DB":       dbName,
			"POSTGRES_USER":     dbUser,
			"POSTGRES_PASSWORD": dbPassword,
		},
		WaitingFor: wait.ForHealthCheckCmd("pg_isready").
			WithRetries(3).
			WithStartPeriod(30 * time.Second).
			WithTestTimeout(5 * time.Second).
			WithTestInterval(10 * time.Second),
	})

	return err
}

```

## Donation

If this project help you reduce time to develop, you can give me a cup of coffee :)

### Paypal donation

[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/donate/?hosted_button_id=PJZSGJN57TDJY)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;or scan this

<img src="https://user-images.githubusercontent.com/1154587/113494222-ad8cb200-94e6-11eb-9ef3-eb883ada222a.png" width="147px" />
