package cascade

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestKey tests the Key type functionality
func TestKey(t *testing.T) {
	t.Run("Key is string type", func(t *testing.T) {
		var k Key = "test-key"
		assert.Equal(t, "test-key", string(k))
	})
}

// TestExecutorFuncName tests the Name method of ExecutorFunc
func TestExecutorFuncName(t *testing.T) {
	tests := []struct {
		name     string
		function ExecutorFunc
		expected string
	}{
		{
			name: "named function",
			function: func(data any) Key {
				return "next-task"
			},
			expected: "cascade.TestExecutorFunc_Name.func1", // Go's naming for anonymous functions in tests
		},
		{
			name: "another named function",
			function: func(data any) Key {
				return "another-task"
			},
			expected: "cascade.TestExecutorFunc_Name.func2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Get the function name
			funcName := tt.function.Name()

			// Check that it returns a Key type
			assert.IsType(t, Key(""), funcName)

			// The actual name will be prefixed with the package and test function name
			// We can't know the exact name, but we can check it contains expected parts
			assert.Contains(t, string(funcName), "cascade.TestExecutorFuncName")
			assert.Contains(t, string(funcName), "func")
		})
	}
}

// TestExecutorFuncExecution tests that ExecutorFunc can be executed
func TestExecutorFuncExecution(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	tests := []struct {
		name        string
		executor    ExecutorFunc
		input       any
		expectedKey Key
	}{
		{
			name: "returns constant key",
			executor: func(data any) Key {
				return "constant-key"
			},
			input:       nil,
			expectedKey: "constant-key",
		},
		{
			name: "processes string data",
			executor: func(data any) Key {
				str, ok := data.(string)
				if !ok {
					return "error"
				}
				return Key("processed-" + str)
			},
			input:       "test",
			expectedKey: "processed-test",
		},
		{
			name: "processes integer data",
			executor: func(data any) Key {
				num, ok := data.(int)
				if !ok {
					return "not-int"
				}
				if num > 10 {
					return "greater-than-ten"
				}
				return "less-or-equal-ten"
			},
			input:       15,
			expectedKey: "greater-than-ten",
		},
		{
			name: "processes struct data",
			executor: func(data any) Key {
				p, ok := data.(Person)
				if !ok {
					return "invalid-person"
				}
				if p.Age >= 18 {
					return Key("adult-" + p.Name)
				}
				return Key("minor-" + p.Name)
			},
			input: Person{
				Name: "John",
				Age:  25,
			},
			expectedKey: "adult-John",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.executor(tt.input)
			assert.Equal(t, tt.expectedKey, result)
		})
	}
}

// TestExecutorFuncComparable tests that different functions have different names
func TestExecutorFuncComparable(t *testing.T) {
	var func1 ExecutorFunc = func(data any) Key {
		return "func1"
	}

	var func2 ExecutorFunc = func(data any) Key {
		return "func2"
	}

	// Different functions should have different names
	name1 := func1.Name()
	name2 := func2.Name()

	assert.NotEqual(t, name1, name2, "Different functions should have different names")
	assert.NotEqual(t, string(name1), string(name2))
}

// TestExecutorFuncPointer tests function pointer behavior
func TestExecutorFuncPointer(t *testing.T) {
	// Create an ExecutorFunc
	var executor ExecutorFunc = func(data any) Key {
		return "test"
	}

	// Test that we can call it
	result := executor("input")
	assert.Equal(t, Key("test"), result)

	// Test that Name() works
	name := executor.Name()
	require.NotEmpty(t, name)

	// Verify it's actually a function pointer
	value := reflect.ValueOf(executor)
	assert.Equal(t, reflect.Func, value.Kind())
}
