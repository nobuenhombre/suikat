package refavour

import (
	"reflect"
)

type TagProcessor interface {
	GetFieldInfo(typeField reflect.StructField, valueField reflect.Value) (interface{}, error)
}

// Получить значение рефлекса
//-------------------------------
func GetReflectValue(s interface{}) reflect.Value {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}

	return val
}

// Проверить Kind
//---------------
func CheckKind(value interface{}, expectedKind reflect.Kind) error {
	reflectValue := GetReflectValue(value)
	reflectKind := reflectValue.Kind()

	if reflectKind != expectedKind {
		return &KindNotMatchedError{
			Expected: expectedKind.String(),
			Actual:   reflectKind.String(),
		}
	}

	return nil
}

// Проверить Структура Ли ?
//-------------------------
func CheckStructure(data interface{}) error {
	return CheckKind(data, reflect.Struct)
}

// Проверить Мапа Ли ?
//-------------------------
func CheckMap(data interface{}) error {
	return CheckKind(data, reflect.Map)
}

// Проверить Слайс Ли ?
//-------------------------
func CheckSlice(data interface{}) error {
	return CheckKind(data, reflect.Slice)
}

// Проверить Можно ли изменять приемник данных
//--------------------------------------------
func CheckCanBeChanged(value interface{}) error {
	reflectValue := GetReflectValue(value)
	if !reflectValue.CanSet() {
		return &CantBeSetError{}
	}

	return nil
}

// Типы полей структуры
//-------------------------------
type FieldsInfo map[string]interface{}

// Функция читает из интерфейса структуру и формирует список имен полей и их типы
//-------------------------------------------------------------------------------
func GetStructureFieldsTypes(value interface{}, tagProcessor TagProcessor) (FieldsInfo, error) {
	reflectValue := GetReflectValue(value)
	reflectedValueType := reflectValue.Type()

	result := make(FieldsInfo)

	for i := 0; i < reflectValue.NumField(); i++ {
		valueField := reflectValue.Field(i)
		typeField := reflectedValueType.Field(i)

		fieldInfo, err := tagProcessor.GetFieldInfo(typeField, valueField)
		if err != nil {
			return nil, err
		}

		result[typeField.Name] = fieldInfo
	}

	return result, nil
}
