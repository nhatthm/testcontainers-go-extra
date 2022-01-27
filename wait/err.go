package wait

// ErrMaxRetriesExceeded indicates that the number of max retries exceeded.
const ErrMaxRetriesExceeded healthCheckError = "max retries exceeded"

type healthCheckError string

// Error satisfies error interface.
func (e healthCheckError) Error() string {
	return string(e)
}
