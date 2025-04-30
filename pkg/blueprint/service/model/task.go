package model

import (
	domainTask "github.com/k4ji/tracesimulator/pkg/model/task"
	"time"
)

// Task represents an operation that can be performed by a service
type Task struct {
	Name                string
	ExternalID          *domainTask.ExternalID
	StartAfter          time.Duration
	Duration            time.Duration
	Kind                string
	Attributes          map[string]string
	Children            []Task
	ChildOf             *domainTask.ExternalID
	LinkedTo            []*domainTask.ExternalID
	FailWithProbability float64
}

// ToRootNodeWithResource converts the Task to a root node with the given resource
func (t *Task) ToRootNodeWithResource(resource *domainTask.Resource) (*domainTask.TreeNode, error) {
	node := domainTask.NewTreeNode(
		domainTask.NewDefinition(
			t.Name,
			true,
			resource,
			t.Attributes,
			domainTask.FromString(t.Kind),
			t.ExternalID,
			t.StartAfter,
			t.Duration,
			t.ChildOf,
			t.LinkedTo,
			t.FailWithProbability,
		),
	)
	for _, child := range t.Children {
		childNode, err := child.toChildNodeWithResource(resource)
		if err != nil {
			return nil, err
		}
		if err = node.AddChild(childNode); err != nil {
			return nil, err
		}
	}
	return node, nil
}

func (t *Task) toChildNodeWithResource(resource *domainTask.Resource) (*domainTask.TreeNode, error) {
	node := domainTask.NewTreeNode(
		domainTask.NewDefinition(
			t.Name,
			false,
			resource,
			t.Attributes,
			domainTask.FromString(t.Kind),
			t.ExternalID,
			t.StartAfter,
			t.Duration,
			nil,
			t.LinkedTo,
			t.FailWithProbability,
		),
	)
	for _, child := range t.Children {
		childNode, err := child.toChildNodeWithResource(resource)
		if err != nil {
			return nil, err
		}
		if err = node.AddChild(childNode); err != nil {
			return nil, err
		}
	}
	return node, nil
}
