package cascade

import (
	"errors"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

var (
	// ErrExecutorNotProvided indicates that an executor function was not provided when creating a task.
	ErrExecutorNotProvided = errors.New("executor not provided")

	// ErrNextTaskNotProvided indicates that a nil next task was provided when linking tasks.
	ErrNextTaskNotProvided = errors.New("next task not provided")

	// ErrNextTaskNotFound indicates that a referenced next task was not found in the task's next map.
	ErrNextTaskNotFound = errors.New("next task not found")
)

// ITask defines the interface for cascade task operations.
// It provides methods for task identification, linking subsequent tasks, and execution.
type ITask interface {
	// Key returns the unique identifier of the task.
	Key() Key

	// LinkNext establishes a connection to a subsequent task in the cascade.
	// Returns an error if the next task is nil.
	LinkNext(*Task) error

	// Run executes the current task and traverses to the next task based on the executor's result.
	// Returns an error if task execution fails or if the next task cannot be found.
	Run() error
}

// Compile-time interface assertion ensures Task implements ITask.
// This validation guarantees that the Task struct satisfies all methods
// defined in the ITask interface. If any method is missing,
// the code will fail to compile, providing early detection of interface compliance issues.
var _ ITask = &Task{}

// Task represents a single unit in a cascade execution flow.
// It contains an executor function, associated data, and references to potential next tasks.
// The cascade execution continues based on the key returned by the executor function.
type Task struct {
	key  Key           // Unique identifier for the task
	data any           // Data passed to the executor function
	exec ExecutorFunc  // Function that executes the task logic and determines the next step
	next map[Key]*Task // Map of possible next tasks keyed by their identifiers
}

// NewTask creates a new Task instance with the provided data and executor function.
// It validates that the executor function is not nil before creating the task.
//
// Parameters:
//   - data: The data to be passed to the executor function
//   - exec: The executor function that defines the task's behavior
//
// Returns:
//   - *Task: A new task instance
//   - error: Returns ErrExecutorNotProvided if the executor function is nil
func NewTask(data any, exec ExecutorFunc) (*Task, error) {
	if exec == nil {
		return nil, ge.Pin(ErrExecutorNotProvided)
	}

	return &Task{
		key:  exec.Name(),
		data: data,
		exec: exec,
		next: make(map[Key]*Task),
	}, nil
}

// Key returns the unique identifier of the task.
// The key is derived from the executor function's Name() method.
//
// Returns:
//   - Key: The task's unique identifier
func (r *Task) Key() Key {
	return r.key
}

// LinkNext establishes a connection from the current task to a subsequent task.
// The linked task will be executed when the executor function returns the corresponding key.
//
// Parameters:
//   - next: The task to be linked as a possible next step
//
// Returns:
//   - error: Returns ErrNextTaskNotProvided if the next task is nil
func (r *Task) LinkNext(next *Task) error {
	if next == nil {
		return ge.Pin(ErrNextTaskNotProvided)
	}

	r.next[next.Key()] = next

	return nil
}

// Run executes the current task and traverses to the next task based on the executor's result.
// The executor function is called with the task's data and returns a key that determines
// which next task to execute. Execution continues until ExitKey is returned or an error occurs.
//
// Returns:
//   - error: Returns ErrNextTaskNotFound if the executor returns a key not found in next map,
//     or propagates any error from the next task's execution
func (r *Task) Run() error {
	// Execute Current Task
	nextKey := r.exec(r.data)

	if nextKey == ExitKey {
		return nil
	}

	// Find Next Task
	nextTask, found := r.next[nextKey]
	if !found {
		return ge.Pin(ErrNextTaskNotFound, ge.Params{"nextKey": nextKey})
	}

	// Call Run for next Task
	err := nextTask.Run()
	if err != nil {
		return ge.Pin(err)
	}

	return nil
}
