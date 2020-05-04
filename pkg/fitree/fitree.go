package fitree

import (
	"os"
)

const StartDepth = 1

// Список Узлов Дерева
type TreeNodeListStruct struct {
	List    []TreeNodeStruct
	Reverse map[string]int
}

// Добавляем Узел
func (list *TreeNodeListStruct) Add(node TreeNodeStruct) {
	list.List = append(list.List, node)
	if list.Reverse == nil {
		list.Reverse = make(map[string]int)
	}
	list.Reverse[node.Path] = len(list.List) - 1
}

// Сканируем Дерево Каталогов
func (list *TreeNodeListStruct) Scan(path string, depth int, ignoreErr bool) error {
	path += string(os.PathSeparator)

	node := TreeNodeStruct{}
	fillErr := node.Fill(path, depth)
	if fillErr != nil && !ignoreErr {
		return fillErr
	}

	list.Add(node)

	depth++

	for _, f := range node.SubDirs {
		scanErr := list.Scan(path+f.Name(), depth, ignoreErr)
		if scanErr != nil && !ignoreErr {
			return scanErr
		}
	}

	return nil
}

// Получить ноду по индексу
func (list *TreeNodeListStruct) GetNode(index int) (*TreeNodeStruct, error) {
	if len(list.List) < index+1 {
		return nil, &NodeIndexDontExistsError{Index: index}
	}

	return &list.List[index], nil
}
