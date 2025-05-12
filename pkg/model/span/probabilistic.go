package span

var _ Condition = (*Probabilistic)(nil)

// Probabilistic triggers based on a random value.
type Probabilistic struct {
	// threshold is the probability threshold for the condition to be met.
	threshold float64
}

func (p Probabilistic) Evaluate(ctx EvaluationContext) bool {
	if ctx.randomness == nil {
		return false
	}
	return ctx.randomness() < p.threshold
}
