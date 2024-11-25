package executiontree

import (
	"errors"
)

type ExecutorNotSetError struct {
	Key string
}

func (e *ExecutorNotSetError) Error() string {
	return "executor is not set"
}

func (e *ExecutorNotSetError) Is(target error) bool {
	var val *ExecutorNotSetError
	if !errors.As(target, &val) {
		return false
	}

	return true
}
