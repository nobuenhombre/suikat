package csvarser

import (
	"reflect"

	"github.com/nobuenhombre/suikat/pkg/refavour"
)

type ParserFunc func(s string) (interface{}, error)

type CsvParser struct {
	TypeParsers map[string]ParserFunc
}

func (p *CsvParser) Init() {
	p.TypeParsers = make(map[string]ParserFunc)
}

func (p *CsvParser) AddTypeParser(dataType string, parser ParserFunc) {
	p.TypeParsers[dataType] = parser
}

// Функция записывает в поле структуры данные
// value - структура
// fieldName - имя поля
// fieldType - тип поля
// data - Данные
//-------------------------------------------
func (p *CsvParser) setStructureFieldData(value interface{}, fieldName string, fieldType reflect.Type, data string) error {
	reflectValue := refavour.GetReflectValue(value)

	parserFunc, found := p.TypeParsers[fieldType.String()]
	if !found {
		return &ParserNotFoundError{
			FieldType: fieldType.String(),
		}
	}
	parsedData, parserErr := parserFunc(data)
	reflectValue.FieldByName(fieldName).Set(reflect.ValueOf(parsedData))
	if parserErr != nil {
		return parserErr
	}

	return nil
}

// Заполнить структуру из слайса
//------------------------------
func (p *CsvParser) FillStructFromSlice(structData interface{}, sliceData []string) error {
	structError := refavour.CheckStructure(structData)
	if structError != nil {
		return structError
	}

	canBeChangedError := refavour.CheckCanBeChanged(structData)
	if canBeChangedError != nil {
		return canBeChangedError
	}

	structureFields, getStructErr := refavour.GetStructureFieldsTypes(structData)
	if getStructErr != nil {
		return getStructErr
	}

	for fieldName, fieldInfo := range structureFields {
		exists := len(sliceData) > fieldInfo.Order
		if exists {
			value := sliceData[fieldInfo.Order]
			setFieldError := p.setStructureFieldData(structData, fieldName, fieldInfo.Type, value)
			if setFieldError != nil {
				return setFieldError
			}
		} else {
			return &FieldNotExistsInSliceError{
				FieldName: fieldName,
				FieldType: fieldInfo.Type.String(),
				Index:     fieldInfo.Order,
			}
		}
	}

	return nil
}
