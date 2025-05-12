package span

import (
	"fmt"
	"github.com/k4ji/tracesimulator/pkg/model/task"
)

// Effect is an interface for applying effects to a tree node.
type Effect interface {
	// Apply applies the effect to the given tree node.
	Apply(node *TreeNode) error
}

// FromEffectSpec converts a task effect specification to an Effect..
func FromEffectSpec(spec task.Effect) (Effect, error) {
	switch spec.Kind() {
	case task.EffectKindMarkAsFailed:
		return &MarkAsFailedEffect{}, nil
	default:
		return nil, fmt.Errorf("unknown effect type: %s", spec.Kind())
	}
}
