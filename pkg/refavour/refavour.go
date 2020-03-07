package refavour

import (
	"reflect"
	"strconv"
)

const TagOrder = "order"

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

type FieldInfo struct {
	Type  reflect.Type
	Order int
}

// Типы полей структуры
//-------------------------------
type FieldsInfo map[string]FieldInfo

// Функция читает из интерфейса структуру и формирует список имен полей и их типы
//-------------------------------------------------------------------------------
func GetStructureFieldsTypes(value interface{}) (FieldsInfo, error) {
	reflectValue := GetReflectValue(value)
	reflectedValueType := reflectValue.Type()

	result := make(FieldsInfo)

	for i := 0; i < reflectValue.NumField(); i++ {
		valueField := reflectValue.Field(i)
		typeField := reflectedValueType.Field(i)

		order, orderErr := strconv.Atoi(typeField.Tag.Get(TagOrder))
		if orderErr != nil {
			return nil, orderErr
		}

		result[typeField.Name] = FieldInfo{
			Type:  valueField.Type(),
			Order: order,
		}
	}
	return result, nil
}
