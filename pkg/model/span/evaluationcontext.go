package span

// EvaluationContext represents the evaluation context for a task.
type EvaluationContext struct {
	randomness func() float64
	node       *TreeNode
}

// NewEvaluationContext creates a new EvaluationContext with the given randomness function and node.
func NewEvaluationContext(randomness func() float64, node *TreeNode) *EvaluationContext {
	return &EvaluationContext{
		randomness: randomness,
		node:       node,
	}
}

// Randomness returns the randomness function for the evaluation context.
func (ec *EvaluationContext) Randomness() func() float64 {
	return ec.randomness
}

// Node returns the node for the evaluation context.
func (ec *EvaluationContext) Node() *TreeNode {
	return ec.node
}
