package span

import "github.com/k4ji/tracesimulator/pkg/model/task"

type AnnotateEffect struct {
	// attributes is a map of attributes to be added to the span.
	attributes map[string]string
}

func (a AnnotateEffect) Apply(node *TreeNode) error {
	for k, v := range a.attributes {
		node.attributes[k] = v
	}
	return nil
}

// FromTaskAnnotateEffect converts a task AnnotateEffect to a span AnnotateEffect.
func FromTaskAnnotateEffect(spec task.AnnotateEffect) AnnotateEffect {
	return AnnotateEffect{
		attributes: spec.Attributes(),
	}
}
