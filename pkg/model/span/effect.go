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
	case task.EffectKindRecordEvent:
		if spec.RecordEventEffect() == nil {
			return nil, fmt.Errorf("record event effect is nil")
		}
		eff, err := FromRecordEventEffect(*spec.RecordEventEffect())
		if err != nil {
			return nil, fmt.Errorf("failed to convert record event effect: %w", err)
		}
		return eff, nil
	default:
		return nil, fmt.Errorf("unknown effect type: %s", spec.Kind())
	}
}
