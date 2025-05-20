package span

var _ Effect = (*DropChildrenEffect)(nil)

// DropChildrenEffect is a conditional definition effect that drops all children of the span.
type DropChildrenEffect struct{}

func (d DropChildrenEffect) Apply(node *TreeNode) error {
	node.children = make([]*TreeNode, 0)
	return nil
}

// NewDropChildrenEffect creates a new DropChildrenEffect.
func NewDropChildrenEffect() DropChildrenEffect {
	return DropChildrenEffect{}
}
