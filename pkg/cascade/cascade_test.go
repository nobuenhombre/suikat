package cascade

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// MockExecutorFunc creates a mock executor function for testing
func MockExecutorFunc(returnKey Key) ExecutorFunc {
	return func(data any) Key {
		return returnKey
	}
}

// MockExecutorWithDataCheck creates a mock function that checks data
func MockExecutorWithDataCheck(t *testing.T, expectedData any, returnKey Key) ExecutorFunc {
	return func(data any) Key {
		assert.Equal(t, expectedData, data)
		return returnKey
	}
}

// MockExecutorWithCallCounter creates a mock function with a call counter
func MockExecutorWithCallCounter(counter *int, returnKey Key) ExecutorFunc {
	return func(data any) Key {
		*counter++
		return returnKey
	}
}

func TestNewTask(t *testing.T) {
	t.Run("Creating a task with a valid executor", func(t *testing.T) {
		exec := MockExecutorFunc("next")
		data := "test data"

		task, err := NewTask(data, exec)

		require.NoError(t, err)
		require.NotNil(t, task)
		assert.Equal(t, exec.Name(), task.Key())
		assert.Equal(t, data, task.data)
		assert.Equal(t, reflect.ValueOf(exec).Pointer(), reflect.ValueOf(task.exec).Pointer())
		assert.NotNil(t, task.next)
		assert.Empty(t, task.next)
	})

	t.Run("Creating a task with nil executor returns an error", func(t *testing.T) {
		task, err := NewTask("data", nil)

		assert.Error(t, err)
		assert.Nil(t, task)

		assert.True(t, assert.ErrorIs(t, err, ErrExecutorNotProvided), "expected error ErrExecutorNotProvided, got: %v", err)
	})

	t.Run("Creating a task with different data types", func(t *testing.T) {
		testCases := []struct {
			name string
			data any
		}{
			{"string", "test string"},
			{"int", 123},
			{"map", map[string]int{"a": 1}},
			{"slice", []string{"a", "b", "c"}},
			{"struct", struct{ Name string }{Name: "test"}},
			{"nil", nil},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				exec := MockExecutorFunc("next")

				task, err := NewTask(tc.data, exec)

				require.NoError(t, err)
				assert.Equal(t, tc.data, task.data)
			})
		}
	})
}

func TestTaskKey(t *testing.T) {
	t.Run("Task key matches executor name", func(t *testing.T) {
		exec := MockExecutorFunc("next")
		task := &Task{
			key:  exec.Name(),
			exec: exec,
		}

		key := task.Key()

		assert.Equal(t, exec.Name(), key)
	})

	t.Run("Task key is unique for different executors", func(t *testing.T) {
		exec1 := MockExecutorFunc("next1")
		exec2 := MockExecutorFunc("next2")

		task1, _ := NewTask(nil, exec1)
		task2, _ := NewTask(nil, exec2)

		assert.NotEqual(t, task1.Key(), task2.Key())
	})
}

func TestTaskLinkNext(t *testing.T) {
	t.Run("Successful task linking", func(t *testing.T) {
		exec1 := MockExecutorFunc("next")
		exec2 := MockExecutorFunc("end")

		task1, _ := NewTask(nil, exec1)
		task2, _ := NewTask(nil, exec2)

		err := task1.LinkNext(task2)

		require.NoError(t, err)
		assert.Len(t, task1.next, 1)
		assert.Equal(t, task2, task1.next[task2.Key()])
	})

	t.Run("Linking multiple next tasks", func(t *testing.T) {
		execMain := MockExecutorFunc("choice")
		execA := MockExecutorFunc("a")
		execB := MockExecutorFunc("b")

		mainTask, _ := NewTask(nil, execMain)
		taskA, _ := NewTask(nil, execA)
		taskB, _ := NewTask(nil, execB)

		err1 := mainTask.LinkNext(taskA)
		err2 := mainTask.LinkNext(taskB)

		require.NoError(t, err1)
		require.NoError(t, err2)
		assert.Len(t, mainTask.next, 2)
		assert.Equal(t, taskA, mainTask.next[taskA.Key()])
		assert.Equal(t, taskB, mainTask.next[taskB.Key()])
	})

	t.Run("Linking with nil task returns an error", func(t *testing.T) {
		task, _ := NewTask(nil, MockExecutorFunc("test"))

		err := task.LinkNext(nil)

		assert.Error(t, err)
		assert.True(t, assert.ErrorIs(t, err, ErrNextTaskNotProvided), "expected error ErrNextTaskNotProvided, got: %v", err)
		assert.Empty(t, task.next)
	})

	t.Run("Overwriting an existing next task", func(t *testing.T) {
		execMain := MockExecutorFunc("next")
		execOld := MockExecutorFunc("old")
		execNew := MockExecutorFunc("new")

		mainTask, _ := NewTask(nil, execMain)
		oldTask, _ := NewTask(nil, execOld)
		newTask, _ := NewTask(nil, execNew)

		// Ensure keys are different
		assert.NotEqual(t, oldTask.Key(), newTask.Key())

		err1 := mainTask.LinkNext(oldTask)
		err2 := mainTask.LinkNext(newTask)

		require.NoError(t, err1)
		require.NoError(t, err2)
		assert.Len(t, mainTask.next, 2)
		assert.Equal(t, oldTask, mainTask.next[oldTask.Key()])
		assert.Equal(t, newTask, mainTask.next[newTask.Key()])
	})
}

func TestTaskRun(t *testing.T) {
	t.Run("Execution termination on ExitKey", func(t *testing.T) {
		counter := 0
		exec := MockExecutorWithCallCounter(&counter, ExitKey)
		task, _ := NewTask("data", exec)

		err := task.Run()

		assert.NoError(t, err)
		assert.Equal(t, 1, counter)
	})

	t.Run("Transition to the next task", func(t *testing.T) {
		callCounters := []int{0, 0}

		exec2 := MockExecutorWithCallCounter(&callCounters[1], ExitKey)
		exec1 := MockExecutorWithCallCounter(&callCounters[0], exec2.Name())

		task1, _ := NewTask("data1", exec1)
		task2, _ := NewTask("data2", exec2)

		err := task1.LinkNext(task2)
		require.NoError(t, err)

		err = task1.Run()

		assert.NoError(t, err)
		assert.Equal(t, 1, callCounters[0])
		assert.Equal(t, 1, callCounters[1])
	})

	t.Run("Multi-level task chain", func(t *testing.T) {
		callCounters := []int{0, 0, 0, 0}

		exec4 := MockExecutorWithCallCounter(&callCounters[3], ExitKey)
		exec3 := MockExecutorWithCallCounter(&callCounters[2], exec4.Name())
		exec2 := MockExecutorWithCallCounter(&callCounters[1], exec3.Name())
		exec1 := MockExecutorWithCallCounter(&callCounters[0], exec2.Name())

		task1, _ := NewTask("data1", exec1)
		task2, _ := NewTask("data2", exec2)
		task3, _ := NewTask("data3", exec3)
		task4, _ := NewTask("data4", exec4)

		task1.LinkNext(task2)
		task2.LinkNext(task3)
		task3.LinkNext(task4)

		err := task1.Run()

		assert.NoError(t, err)
		assert.Equal(t, []int{1, 1, 1, 1}, callCounters)
	})

	t.Run("Execution branching", func(t *testing.T) {
		counterMain := 0
		counterA := 0
		counterB := 0

		execA := MockExecutorWithCallCounter(&counterA, ExitKey)
		execB := MockExecutorWithCallCounter(&counterB, ExitKey)

		execMain := func(data any) Key {
			counterMain++
			// Return "taskA" or "taskB" depending on data
			if data == "goA" {
				return execA.Name()
			}

			return execB.Name()
		}

		mainTask, _ := NewTask("goA", execMain)
		taskA, _ := NewTask(nil, execA)
		taskB, _ := NewTask(nil, execB)

		mainTask.LinkNext(taskA)
		mainTask.LinkNext(taskB)

		err := mainTask.Run()

		assert.NoError(t, err)
		assert.Equal(t, 1, counterMain)
		assert.Equal(t, 1, counterA)
		assert.Equal(t, 0, counterB) // Should not be called
	})

	t.Run("Error when next task is missing", func(t *testing.T) {
		exec := MockExecutorFunc("nonExistentKey")
		task, _ := NewTask(nil, exec)

		err := task.Run()

		assert.Error(t, err)
		assert.True(t, assert.ErrorIs(t, err, ErrNextTaskNotFound), "expected error ErrNextTaskNotFound, got: %v", err)
	})

	t.Run("Recursive Run call for task chain", func(t *testing.T) {
		var func4 ExecutorFunc = func(data any) Key { return ExitKey }
		var func3 ExecutorFunc = func(data any) Key { return func4.Name() }
		var func2 ExecutorFunc = func(data any) Key { return func3.Name() }
		var func1 ExecutorFunc = func(data any) Key { return func2.Name() }

		execChain := []ExecutorFunc{func1, func2, func3, func4}

		tasks := make([]*Task, len(execChain))
		for i, exec := range execChain {
			task, err := NewTask(i, exec)
			require.NoError(t, err)
			tasks[i] = task
		}

		// Link tasks into a chain
		for i := 0; i < len(tasks)-1; i++ {
			err := tasks[i].LinkNext(tasks[i+1])
			require.NoError(t, err)
		}

		err := tasks[0].Run()

		assert.NoError(t, err)
	})
}
