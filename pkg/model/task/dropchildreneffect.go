package task

// DropChildrenEffect is a conditional definition effect that drops all children of the task.
type DropChildrenEffect struct{}

// NewDropChildrenEffect creates a new DropChildrenEffect.
func NewDropChildrenEffect() DropChildrenEffect {
	return DropChildrenEffect{}
}
