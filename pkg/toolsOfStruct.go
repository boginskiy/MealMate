package pkg

import (
	"reflect"
	"strings"
)

type ToolsStructer interface {
	ShowdownFull(any) map[string]map[string]any
	ShowdownFields(any) (fieldsList []string)
}

type ToolsOfStruct struct {
}

func NewToolsOfStruct() *ToolsOfStruct {
	return &ToolsOfStruct{}
}

func (t *ToolsOfStruct) ShowdownFull(someStruct any) map[string]map[string]any {
	infoStruct := map[string]map[string]any{}

	_type := reflect.TypeOf(someStruct)
	_value := reflect.ValueOf(someStruct)

	if _type.Kind() == reflect.Struct {
		numFields := _type.NumField()

		// Проходим по каждому полю структуры
		for i := 0; i < numFields; i++ {
			field := _type.Field(i)
			name := field.Name
			kind := field.Type.Kind()
			value := _value.FieldByName(name).Interface()

			infoStruct[name] = map[string]any{"kind": kind.String(), "value": value, "count": numFields}
		}
	}
	return infoStruct
}

func (t *ToolsOfStruct) ShowdownFields(someStruct any) (fieldsList []string) {
	_type := reflect.TypeOf(someStruct)

	if _type.Kind() == reflect.Struct {
		numFields := _type.NumField()
		fieldsList = make([]string, numFields)

		// Собираем поля структуры
		for i := 0; i < numFields; i++ {
			field := _type.Field(i)
			fieldsList[i] = strings.ToLower(field.Name)
		}
	}
	return fieldsList
}
