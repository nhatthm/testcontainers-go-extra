package testcontainers

import (
	"fmt"

	"github.com/testcontainers/testcontainers-go"
)

// WithNamePrefix sets a prefix for the request name.
func WithNamePrefix(prefix string) GenericContainerOption {
	return genericContainerOptionFunc(func(o *genericContainerOptions) {
		o.request.Name = fmt.Sprintf("%s_%s", prefix, o.request.Name)
	})
}

// WithNameSuffix sets a suffix for the request name.
func WithNameSuffix(suffix string) GenericContainerOption {
	return genericContainerOptionFunc(func(o *genericContainerOptions) {
		o.request.Name = fmt.Sprintf("%s_%s", o.request.Name, suffix)
	})
}

// WithImageName sets the image name.
func WithImageName(name string) GenericContainerOption {
	return genericContainerOptionFunc(func(o *genericContainerOptions) {
		o.request.Image = changeImageName(o.request.Image, name)
	})
}

// WithImageTag sets the image tag.
func WithImageTag(tag string) GenericContainerOption {
	return genericContainerOptionFunc(func(o *genericContainerOptions) {
		o.request.Image = changeImageTag(o.request.Image, tag)
	})
}

// WithCallback adds a new callback to run after the container is ready.
func WithCallback(f ContainerCallback) GenericContainerOption {
	return genericContainerOptionFunc(func(o *genericContainerOptions) {
		o.callbacks = append(o.callbacks, f)
	})
}

// WithProviderType sets the provider type.
func WithProviderType(providerType testcontainers.ProviderType) GenericContainerOption {
	return genericContainerOptionFunc(func(o *genericContainerOptions) {
		o.providerType = providerType
	})
}
