package task_test

import (
	"testing"
	"time"

	"github.com/k4ji/tracesimulator/pkg/model/task"
	"github.com/k4ji/tracesimulator/pkg/model/task/taskduration"
	"github.com/stretchr/testify/assert"
)

func TestNewDelay(t *testing.T) {
	t.Run("valid expression", func(t *testing.T) {
		expr, _ := taskduration.NewAbsoluteDuration(5 * time.Second)
		delay, err := task.NewDelay(expr)

		assert.NoError(t, err)
		assert.NotNil(t, delay)
	})

	t.Run("nil expression", func(t *testing.T) {
		assert.Panics(t, func() {
			_, _ = task.NewDelay(nil)
		}, "Expected panic when expression is nil")
	})
}

func TestDelay_Resolve(t *testing.T) {
	t.Run("resolve absolute duration", func(t *testing.T) {
		expr, _ := taskduration.NewAbsoluteDuration(5 * time.Second)
		delay, _ := task.NewDelay(expr)

		result, err := delay.Resolve(nil)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 5*time.Second, *result)
	})

	t.Run("resolve relative duration with valid context", func(t *testing.T) {
		baseDuration := 10 * time.Second
		expr, _ := taskduration.NewRelativeDuration(0.5)
		delay, _ := task.NewDelay(expr)

		result, err := delay.Resolve(&baseDuration)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 5*time.Second, *result)
	})

	t.Run("resolve relative duration with invalid context type", func(t *testing.T) {
		expr, _ := taskduration.NewRelativeDuration(0.5)
		delay, _ := task.NewDelay(expr)

		result, err := delay.Resolve("invalid context")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "invalid context type")
	})

	t.Run("resolve relative duration with nil context", func(t *testing.T) {
		expr, _ := taskduration.NewRelativeDuration(0.5)
		delay, _ := task.NewDelay(expr)

		result, err := delay.Resolve(nil)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "invalid context type")
	})
}
