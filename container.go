package testcontainers

import (
	"context"
	"fmt"
	"sync"

	"github.com/testcontainers/testcontainers-go"
)

// GenericContainerOption is option for starting a new generic container.
type GenericContainerOption interface {
	applyOptions(o *genericContainerOptions)
}

type genericContainerOptionFunc func(o *genericContainerOptions)

func (f genericContainerOptionFunc) applyOptions(o *genericContainerOptions) {
	f(o)
}

type genericContainerOptions struct {
	request      ContainerRequest
	providerType testcontainers.ProviderType
	callbacks    []ContainerCallback
}

// StartGenericContainer starts a new generic container.
func StartGenericContainer(ctx context.Context, request ContainerRequest, opts ...GenericContainerOption) (Container, error) {
	originalName := request.Name

	o := genericContainerOptions{
		request: request,
	}

	for _, opt := range opts {
		opt.applyOptions(&o)
	}

	r := testcontainers.GenericContainerRequest{
		ContainerRequest: o.request,
		Started:          true,
		ProviderType:     o.providerType,
	}
	o.request.Name = originalName

	c, err := testcontainers.GenericContainer(ctx, r)
	if err != nil {
		return c, err
	}

	for _, f := range o.callbacks {
		if err := f(ctx, c, o.request); err != nil {
			return c, err
		}
	}

	return c, nil
}

// StartGenericContainerRequest is request for starting a new generic container.
type StartGenericContainerRequest struct {
	Request ContainerRequest
	Options []GenericContainerOption
}

// StartGenericContainers starts multiple generic containers at once.
func StartGenericContainers(ctx context.Context, requests ...StartGenericContainerRequest) (containers []Container, _ error) {
	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)

	wg.Add(len(requests))

	containers = make([]Container, 0, len(requests))
	errs := make(errorCollection, 0, len(requests))

	for _, r := range requests {
		go func(r StartGenericContainerRequest) {
			defer wg.Done()

			c, sErr := StartGenericContainer(ctx, r.Request, r.Options...)

			mu.Lock()
			defer mu.Unlock()

			if sErr != nil {
				errs.Append(fmt.Errorf("could not start container %q: %w", r.Request.Name, sErr))
			}

			if c != nil {
				containers = append(containers, c)
			}
		}(r)
	}

	wg.Wait()

	return containers, errs.AsError()
}

// StopGenericContainers stops multiple containers at once.
func StopGenericContainers(ctx context.Context, containers ...Container) error {
	var wg sync.WaitGroup

	wg.Add(len(containers))

	errs := make(errorCollection, 0, len(containers))

	for _, c := range containers {
		go func(c testcontainers.Container) {
			defer wg.Done()

			if err := c.Terminate(ctx); err != nil {
				errs.Append(fmt.Errorf("could not stop container %q: %w", c.GetContainerID(), err))
			}
		}(c)
	}

	wg.Wait()

	return errs.AsError()
}

// ContainerCallback is called after a container is successfully created and started.
type ContainerCallback func(ctx context.Context, c Container, r ContainerRequest) error

func (c ContainerCallback) applyOptions(o *genericContainerOptions) {
	o.callbacks = append(o.callbacks, c)
}
