package testcontainers

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustNotFail(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		mustNotFail(nil)
	})

	assert.Panics(t, func() {
		mustNotFail(errors.New("error"))
	})
}

func TestJoinStrings(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario string
		prefix   string
		name     string
		suffix   string
		expected string
	}{
		{
			scenario: "no prefix or suffix",
			name:     "container",
			expected: "container",
		},
		{
			scenario: "only prefix",
			prefix:   "prefix",
			name:     "container",
			expected: "prefix_container",
		},
		{
			scenario: "only suffix",
			name:     "container",
			suffix:   "suffix",
			expected: "container_suffix",
		},
		{
			scenario: "both prefix and suffix",
			prefix:   "prefix",
			name:     "container",
			suffix:   "suffix",
			expected: "prefix_container_suffix",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual := joinStrings(tc.prefix, tc.name, tc.suffix)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestChangeImageName(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario string
		image    string
		name     string
		expected string
	}{
		{
			scenario: "no domain and no tag",
			image:    "mysql",
			name:     "percona",
			expected: "percona",
		},
		{
			scenario: "no domain and with tag",
			image:    "mysql:8",
			name:     "percona",
			expected: "percona:8",
		},
		{
			scenario: "with path and no tag",
			image:    "org/mysql:8",
			name:     "percona",
			expected: "percona:8",
		},
		{
			scenario: "with path and tag",
			image:    "org/mysql:8",
			name:     "percona",
			expected: "percona:8",
		},
		{
			scenario: "with domain and no tag",
			image:    "ghcr.io/org/mysql",
			name:     "percona",
			expected: "ghcr.io/percona",
		},
		{
			scenario: "with domain and tag",
			image:    "ghcr.io/org/mysql:8",
			name:     "percona",
			expected: "ghcr.io/percona:8",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual := changeImageName(tc.image, tc.name)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestChangeImageTag(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario string
		image    string
		tag      string
		expected string
	}{
		{
			scenario: "no tag",
			image:    "mysql",
			tag:      "8",
			expected: "mysql:8",
		},
		{
			scenario: "with tag",
			image:    "mysql:8",
			tag:      "9",
			expected: "mysql:9",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			actual := changeImageTag(tc.image, tc.tag)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
