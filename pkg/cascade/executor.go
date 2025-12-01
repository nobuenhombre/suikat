package cascade

import (
	"reflect"
	"runtime"
)

// Key represents a string identifier for an executor function.
// It's used to uniquely identify executor functions within the cascade system.
type Key string

// ExitKey is a predefined key used to signal the termination of task execution flow.
const ExitKey Key = "exit"

// ExecutorFunc defines the function signature for task execution.
// It takes task data as input and returns the Key of the next task to execute.
type ExecutorFunc func(data any) Key

// Name returns the name of the executor function as an ExecutorFuncName.
// This method extracts the full function name (including package path) using reflection.
//
// Returns:
//   - ExecutorFuncName: The full name of the function including package path
func (f ExecutorFunc) Name() Key {
	// Get the full function name (including package)
	fullName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()

	return Key(fullName)
}
