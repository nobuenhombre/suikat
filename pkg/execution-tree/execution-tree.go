package executiontree

import "github.com/nobuenhombre/suikat/pkg/ge"

type NodeKey string

const NodeKeyExit = "exit"

type NodeFunc func(input interface{}) NodeKey

type Node struct {
	Executor      NodeFunc
	ExecutorInput interface{}
	Branches      map[NodeKey]*Node
}

type IExecutionTree interface {
	AddBranch(key NodeKey, branchNode *Node)
	Run() error
}

func NewNode(executor NodeFunc, executorInput interface{}) *Node {
	return &Node{
		Executor:      executor,
		ExecutorInput: executorInput,
		Branches:      make(map[NodeKey]*Node),
	}
}

func (node *Node) AddBranch(key NodeKey, branchNode *Node) {
	node.Branches[key] = branchNode
}

func (node *Node) Run() error {
	if node.Executor == nil {
		return ge.Pin(&ExecutorNotSetError{})
	}

	key := node.Executor(node.ExecutorInput)

	if key == NodeKeyExit {
		return nil
	}

	nextNode, ok := node.Branches[key]
	if !ok {
		return ge.Pin(&ge.NotFoundError{
			Key: string(key),
		})
	}

	err := nextNode.Run()
	if err != nil {
		return ge.Pin(err)
	}

	return nil
}
