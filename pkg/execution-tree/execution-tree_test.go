package executiontree

import (
	"github.com/nobuenhombre/suikat/pkg/ge"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNilExecutor(t *testing.T) {
	node := NewNode(nil, nil)

	err := node.Run()
	assert.Error(t, err)
	assert.ErrorIs(t, err, &ExecutorNotSetError{})
}

func TestExitNode(t *testing.T) {
	executor := func(input interface{}) NodeKey {
		return NodeKeyExit
	}

	node := NewNode(executor, nil)

	err := node.Run()
	assert.NoError(t, err)
}

func TestNotFoundNodeError(t *testing.T) {
	executorA := func(input interface{}) NodeKey {
		return "B"
	}

	executorB := func(input interface{}) NodeKey {
		return "D"
	}

	nodeA := NewNode(executorA, nil)
	nodeB := NewNode(executorB, nil)

	nodeA.AddBranch("B", nodeB)

	err := nodeA.Run()
	assert.ErrorIs(t, err, &ge.NotFoundError{Key: "D"})
}

func TestExecutorTree(t *testing.T) {
	type Sum int
	var TestSum Sum

	executorA := func(input interface{}) NodeKey {
		sum := input.(*Sum)
		*sum += 1
		return "B"
	}

	executorB := func(input interface{}) NodeKey {
		sum := input.(*Sum)
		*sum += 2
		return "C"
	}

	executorC := func(input interface{}) NodeKey {
		sum := input.(*Sum)
		*sum += 3
		return "exit"
	}

	nodeA := NewNode(executorA, &TestSum)
	nodeB := NewNode(executorB, &TestSum)
	nodeC := NewNode(executorC, &TestSum)

	nodeA.AddBranch("B", nodeB)
	nodeB.AddBranch("C", nodeC)

	err := nodeA.Run()
	assert.NoError(t, err)

	assert.Equal(t, Sum(6), TestSum)
}
