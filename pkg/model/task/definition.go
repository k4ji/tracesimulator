package task

import (
	"fmt"
	"github.com/k4ji/tracesimulator/pkg/model/task/taskduration"
	"time"
)

// Definition represents a task in the trace
type Definition struct {
	name                 string
	isResourceEntryPoint bool
	resource             *Resource
	attributes           map[string]string
	kind                 Kind
	externalID           *ExternalID
	delay                taskduration.Expression // Relative time from the start of the parent task
	duration             time.Duration           // Relative time from the start of the parent task
	childOf              *ExternalID             // ID of the parent task (if any)
	linkedTo             []*ExternalID           // IDs of linked spans (for producer/consumer relationships)
	failWithProbability  float64                 // Probability of error
}

// NewDefinition creates a new task definition
func NewDefinition(name string, isResourceEntryPoint bool, resource *Resource, attributes map[string]string, kind Kind, externalID *ExternalID, delay taskduration.Expression, duration time.Duration, childOf *ExternalID, linkedTo []*ExternalID, failWithProbability float64) (*Definition, error) {
	if delay == nil {
		return nil, fmt.Errorf("delay cannot be nil")
	}
	return &Definition{
		name:                 name,
		isResourceEntryPoint: isResourceEntryPoint,
		resource:             resource,
		attributes:           attributes,
		kind:                 kind,
		externalID:           externalID,
		delay:                delay,
		duration:             duration,
		childOf:              childOf,
		linkedTo:             linkedTo,
		failWithProbability:  failWithProbability,
	}, nil
}

func (d *Definition) Name() string {
	return d.name
}

func (d *Definition) Resource() *Resource {
	return d.resource
}

func (d *Definition) Attributes() map[string]string {
	return d.attributes
}

func (d *Definition) IsResourceEntryPoint() bool {
	return d.isResourceEntryPoint
}

func (d *Definition) Kind() Kind {
	return d.kind
}

func (d *Definition) ExternalID() *ExternalID {
	return d.externalID
}

func (d *Definition) Delay() taskduration.Expression {
	return d.delay
}

func (d *Definition) Duration() time.Duration {
	return d.duration
}

func (d *Definition) ChildOf() *ExternalID {
	return d.childOf
}

func (d *Definition) LinkedTo() []*ExternalID {
	return d.linkedTo
}

func (d *Definition) FailWithProbability() float64 {
	return d.failWithProbability
}
