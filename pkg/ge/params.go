package ge

import (
	"fmt"
	"strings"
)

type Params map[string]interface{}

func (p Params) View() string {
	result := ""
	for key, value := range p {
		result += fmt.Sprintf("(%v = %v), ", key, value)
	}

	return strings.TrimSuffix(result, ", ")
}
