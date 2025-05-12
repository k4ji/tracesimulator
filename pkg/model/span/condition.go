package span

import (
	"fmt"
	"github.com/k4ji/tracesimulator/pkg/model/task"
)

// Condition is an interface for evaluating whether a condition is met.
type Condition interface {
	// Evaluate evaluates the condition based on the provided context and returns true if the condition is met.
	Evaluate(ctx EvaluationContext) bool
}

// FromConditionSpec converts a Condition spec to a Condition.
func FromConditionSpec(spec task.Condition) (Condition, error) {
	switch spec.Kind() {
	case task.ConditionKindProbabilistic:
		if spec.Probabilistic() == nil {
			return nil, fmt.Errorf("probabilistic condition requires a probability")
		}
		return Probabilistic{threshold: spec.Probabilistic().Threshold()}, nil
	default:
		return nil, fmt.Errorf("unsupported condition type: %s", spec.Kind())
	}
}
