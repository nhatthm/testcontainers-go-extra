package testcontainers

import (
	"fmt"
)

type errorCollection []error

func (e *errorCollection) Append(err error) {
	*e = append(*e, err)
}

func (e *errorCollection) AsError() error {
	if len(*e) == 0 {
		return nil
	}

	err := (*e)[0]

	for i := 1; i < len(*e); i++ {
		err = fmt.Errorf("%w\n%s", err, (*e)[i].Error())
	}

	return err
}
