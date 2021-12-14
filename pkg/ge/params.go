package ge

import (
	"fmt"
	"strings"
)

// Params
// en: error parameters - variables in the presence of which the error occurred
// - will help you understand what happened
// ru: параметры ошибки - переменные в присутствии которых произошла ошибка - помогут понять что произошло
type Params map[string]interface{}

// View
// en: returns parameters and their values as a string
// ru: возвращает параметры и их значения в виде строки
func (p Params) View() string {
	result := ""
	for key, value := range p {
		result += fmt.Sprintf("(%v = %v), ", key, value)
	}

	return strings.TrimSuffix(result, ", ")
}
