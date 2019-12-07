package application

type Container interface {
	// Register a binding with the container.
	bind()

	// Register a shared binding in the container.
	singleton()
}
