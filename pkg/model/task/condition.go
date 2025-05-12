package task

type ConditionKind string

const (
	ConditionKindProbabilistic ConditionKind = "probabilistic"
)

// Condition is an interface for evaluating whether an effect should be applied.
type Condition struct {
	// kind is the type of condition.
	kind ConditionKind
	// probabilistic is the probability of the condition being met.
	probabilistic *Probabilistic
}

// NewProbabilisticCondition creates a new Condition with the given probability.
func NewProbabilisticCondition(threshold float64) Condition {
	return Condition{
		kind: ConditionKindProbabilistic,
		probabilistic: &Probabilistic{
			threshold: threshold,
		},
	}
}

func (c Condition) Kind() ConditionKind {
	return c.kind
}

func (c Condition) Probabilistic() *Probabilistic {
	return c.probabilistic
}
