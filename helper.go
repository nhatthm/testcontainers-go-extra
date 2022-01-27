package testcontainers

import (
	"strings"

	"github.com/docker/distribution/reference"
)

func mustNotFail(err error) {
	if err != nil {
		panic(err)
	}
}

func joinStrings(prefix, s, suffix string) string {
	var sb strings.Builder

	if prefix != "" {
		sb.WriteString(prefix)
		sb.WriteRune('_')
	}

	sb.WriteString(s)

	if suffix != "" {
		sb.WriteRune('_')
		sb.WriteString(suffix)
	}

	return sb.String()
}

func changeImageName(image string, name string) string {
	named, err := reference.ParseNormalizedNamed(image)
	mustNotFail(err)

	if domain := reference.Domain(named); domain != "" {
		name = domain + "/" + name
	}

	result, err := reference.WithName(name)
	mustNotFail(err)

	if tagged, ok := named.(reference.NamedTagged); ok {
		result, err = reference.WithTag(result, tagged.Tag())
		mustNotFail(err)
	}

	return reference.FamiliarString(result)
}

func changeImageTag(image string, tag string) string {
	named, err := reference.ParseNormalizedNamed(image)
	mustNotFail(err)

	named, err = reference.WithTag(named, tag)
	mustNotFail(err)

	return reference.FamiliarString(named)
}
