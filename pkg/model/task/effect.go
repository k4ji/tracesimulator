package task

// EffectKind defines the type of effect to apply.
type EffectKind string

const (
	EffectKindMarkAsFailed EffectKind = "markAsFailed"
)

type Effect struct {
	kind         EffectKind
	markAsFailed *MarkAsFailedEffect
}

func FromMarkAsFailedEffect(markAsFailed MarkAsFailedEffect) Effect {
	return Effect{
		kind:         EffectKindMarkAsFailed,
		markAsFailed: &markAsFailed,
	}
}

func (e *Effect) Kind() EffectKind {
	return e.kind
}

func (e *Effect) MarkAsFailedEffect() *MarkAsFailedEffect {
	return e.markAsFailed
}
