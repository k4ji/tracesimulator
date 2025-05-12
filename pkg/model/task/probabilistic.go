package task

// Probabilistic represents a condition that triggers based on a random value.
type Probabilistic struct {
	// threshold is the probability threshold for the condition to be met.
	threshold float64
}

// NewProbabilistic creates a new Probabilistic condition with the given threshold.
func NewProbabilistic(threshold float64) Probabilistic {
	return Probabilistic{
		threshold: threshold,
	}
}

// Threshold returns the probability threshold for the condition to be met.
func (p Probabilistic) Threshold() float64 {
	return p.threshold
}
