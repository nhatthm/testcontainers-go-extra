package testcontainers

// Container statuses.
const (
	ContainerStatusCreated    ContainerStatus = "created"
	ContainerStatusRunning    ContainerStatus = "running"
	ContainerStatusPaused     ContainerStatus = "paused"
	ContainerStatusRestarting ContainerStatus = "restarting"
	ContainerStatusRemoving   ContainerStatus = "removing"
	ContainerStatusExited     ContainerStatus = "exited"
	ContainerStatusDead       ContainerStatus = "dead"
)

// ContainerStatus is status of a container.
type ContainerStatus string

// Equal checks if two statuses are the same.
func (s ContainerStatus) Equal(target string) bool {
	return string(s) == target
}
