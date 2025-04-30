package span

import (
	"testing"
	"time"

	"github.com/k4ji/tracesimulator/pkg/model/task"
	"github.com/stretchr/testify/assert"
)

func TestFromTaskTree(t *testing.T) {
	type testCase struct {
		name        string
		taskTree    *task.TreeNode
		traceID     TraceID
		baseEndTime time.Time
		idGen       func() ID
		statusGen   func(prob float64) Status
		expected    *TreeNode
	}

	baseTime := time.Now()
	traceID := NewTraceID([16]byte{0x01})

	testCases := []testCase{
		{
			name: "transform a root task to a span",
			taskTree: task.NewTreeNode(
				task.NewDefinition(
					"root-task",
					true,
					task.NewResource("service-a", map[string]string{"service.version": "1.0.0"}),
					map[string]string{"team": "team-a"},
					task.KindServer,
					func() *task.ExternalID { id, _ := task.NewExternalID("root-task"); return id }(),
					1*time.Second,
					2*time.Second,
					nil,
					[]*task.ExternalID{},
					0.0,
				),
			),
			traceID:     traceID,
			baseEndTime: baseTime,
			idGen:       func() ID { return NewSpanID([8]byte{0x01}) },
			statusGen:   func(prob float64) Status { return StatusOK },
			expected: &TreeNode{
				id:                   NewSpanID([8]byte{0x01}),
				traceID:              traceID,
				name:                 "root-task",
				isResourceEntryPoint: true,
				resource:             task.NewResource("service-a", map[string]string{"service.version": "1.0.0"}),
				attributes:           map[string]string{"team": "team-a"},
				kind:                 KindServer,
				startTime:            baseTime.Add(1 * time.Second),
				endTime:              baseTime.Add(3 * time.Second),
				parentID:             nil,
				externalID:           func() *task.ExternalID { id, _ := task.NewExternalID("root-task"); return id }(),
				status:               StatusOK,
				children:             []*TreeNode{},
				linkedTo:             []*TreeNode{},
				linkedToExternalID:   []*task.ExternalID{},
			},
		},
		{
			name: "copy resource and attributes of a task into a span",
			taskTree: func() *task.TreeNode {
				root := task.NewTreeNode(
					task.NewDefinition(
						"root-task",
						true,
						task.NewResource("service-a", map[string]string{"service.version": "1.0.0"}),
						map[string]string{"key1": "val1"},
						task.KindInternal,
						nil,
						1*time.Second,
						2*time.Second,
						nil,
						[]*task.ExternalID{},
						0.0,
					),
				)
				//nolint:errcheck
				root.AddChild(
					task.NewTreeNode(
						task.NewDefinition(
							"child-task",
							false,
							task.NewResource("service-a", map[string]string{"service.version": "1.0.0"}),
							map[string]string{"key2": "val2"},
							task.KindClient,
							nil,
							3*time.Second,
							4*time.Second,
							nil,
							[]*task.ExternalID{},
							0.0,
						),
					),
				)
				return root
			}(),
			traceID:     traceID,
			baseEndTime: baseTime,
			idGen: func() func() ID {
				ids := [][8]byte{
					{0x01}, // ID for the root span
					{0x02}, // ID for the child span
				}
				index := 0
				return func() ID {
					id := NewSpanID(ids[index])
					index++
					return id
				}
			}(),
			statusGen: func(prob float64) Status { return StatusOK },
			expected: &TreeNode{
				id:                   NewSpanID([8]byte{0x01}),
				traceID:              traceID,
				name:                 "root-task",
				isResourceEntryPoint: true,
				kind:                 KindInternal,
				resource:             task.NewResource("service-a", map[string]string{"service.version": "1.0.0"}),
				attributes:           map[string]string{"key1": "val1"},
				startTime:            baseTime.Add(1 * time.Second),
				endTime:              baseTime.Add(3 * time.Second),
				status:               StatusOK,
				children: []*TreeNode{
					{
						id:                   NewSpanID([8]byte{0x02}),
						traceID:              traceID,
						name:                 "child-task",
						isResourceEntryPoint: false,
						kind:                 KindClient,
						resource:             task.NewResource("service-a", map[string]string{"service.version": "1.0.0"}),
						attributes:           map[string]string{"key2": "val2"},
						startTime:            baseTime.Add(4 * time.Second),
						endTime:              baseTime.Add(8 * time.Second),
						status:               StatusOK,
						parentID:             func() *ID { id := NewSpanID([8]byte{0x01}); return &id }(),
						externalID:           nil,
						linkedTo:             []*TreeNode{},
						linkedToExternalID:   []*task.ExternalID{},
						children:             []*TreeNode{},
					},
				},
				linkedTo:           []*TreeNode{},
				linkedToExternalID: []*task.ExternalID{},
			},
		},
		{
			name: "set start time and duration relative to the parent span",
			taskTree: func() *task.TreeNode {
				root := task.NewTreeNode(
					task.NewDefinition(
						"root-task",
						true,
						task.NewResource("service-a", make(map[string]string)),
						make(map[string]string),
						task.KindInternal,
						nil,
						1*time.Second,
						2*time.Second,
						nil,
						[]*task.ExternalID{},
						0.0,
					),
				)
				//nolint:errcheck
				root.AddChild(
					task.NewTreeNode(
						task.NewDefinition(
							"child-task",
							false,
							task.NewResource("service-a", make(map[string]string)),
							make(map[string]string),
							task.KindClient,
							nil,
							3*time.Second,
							4*time.Second,
							nil,
							[]*task.ExternalID{},
							0.0,
						),
					),
				)
				return root
			}(),
			traceID:     traceID,
			baseEndTime: baseTime,
			idGen: func() func() ID {
				ids := [][8]byte{
					{0x01}, // ID for the root span
					{0x02}, // ID for the child span
				}
				index := 0
				return func() ID {
					id := NewSpanID(ids[index])
					index++
					return id
				}
			}(),
			statusGen: func(prob float64) Status { return StatusOK },
			expected: &TreeNode{
				id:                   NewSpanID([8]byte{0x01}),
				traceID:              traceID,
				name:                 "root-task",
				isResourceEntryPoint: true,
				kind:                 KindInternal,
				resource:             task.NewResource("service-a", make(map[string]string)),
				attributes:           make(map[string]string),
				startTime:            baseTime.Add(1 * time.Second),
				endTime:              baseTime.Add(3 * time.Second),
				status:               StatusOK,
				children: []*TreeNode{
					{
						id:                   NewSpanID([8]byte{0x02}),
						traceID:              traceID,
						name:                 "child-task",
						isResourceEntryPoint: false,
						kind:                 KindClient,
						resource:             task.NewResource("service-a", make(map[string]string)),
						attributes:           make(map[string]string),
						startTime:            baseTime.Add(4 * time.Second),
						endTime:              baseTime.Add(8 * time.Second),
						status:               StatusOK,
						parentID:             func() *ID { id := NewSpanID([8]byte{0x01}); return &id }(),
						externalID:           nil,
						linkedTo:             []*TreeNode{},
						linkedToExternalID:   []*task.ExternalID{},
						children:             []*TreeNode{},
					},
				},
				linkedTo:           []*TreeNode{},
				linkedToExternalID: []*task.ExternalID{},
			},
		},
		{
			name: "generate error spans based on fail probability",
			taskTree: task.NewTreeNode(
				task.NewDefinition(
					"root-task",
					true,
					task.NewResource("service-a", make(map[string]string)),
					make(map[string]string),
					task.KindInternal,
					nil,
					1*time.Second,
					2*time.Second,
					nil,
					[]*task.ExternalID{},
					0.5,
				),
			),
			traceID:     traceID,
			baseEndTime: baseTime,
			idGen:       func() ID { return NewSpanID([8]byte{0x01}) },
			statusGen: func(prob float64) Status {
				if prob > 0 {
					return StatusError
				}
				return StatusOK
			},
			expected: &TreeNode{
				id:                   NewSpanID([8]byte{0x01}),
				traceID:              traceID,
				name:                 "root-task",
				isResourceEntryPoint: true,
				kind:                 KindInternal,
				resource:             task.NewResource("service-a", make(map[string]string)),
				attributes:           make(map[string]string),
				startTime:            baseTime.Add(1 * time.Second),
				endTime:              baseTime.Add(3 * time.Second),
				status:               StatusError,
				linkedTo:             []*TreeNode{},
				linkedToExternalID:   []*task.ExternalID{},
				children:             []*TreeNode{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			span, err := FromTaskTree(tc.taskTree, tc.traceID, tc.baseEndTime, tc.idGen, tc.statusGen)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, span)
		})
	}
}

func TestShiftTimestamps(t *testing.T) {
	now := time.Now()
	rootNodeStartTime := now.Add(0 * time.Second)
	rootNodeEndTime := now.Add(1 * time.Second)
	childNodeStartTime := now.Add(2 * time.Second)
	childNodeEndTime := now.Add(3 * time.Second)
	node := &TreeNode{
		startTime: rootNodeStartTime,
		endTime:   rootNodeEndTime,
		children: []*TreeNode{
			{
				startTime: childNodeStartTime,
				endTime:   childNodeEndTime,
			},
		},
	}

	delta := 3 * time.Second
	node.ShiftTimestamps(delta)

	assert.Equal(t, rootNodeStartTime.Add(delta), node.startTime)
	assert.Equal(t, rootNodeEndTime.Add(delta), node.endTime)
	assert.Equal(t, childNodeStartTime.Add(delta), node.children[0].startTime)
	assert.Equal(t, childNodeEndTime.Add(delta), node.children[0].endTime)
}
