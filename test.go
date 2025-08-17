package main

// TODO
// Тесты надо

import (
	"fmt"
	"reflect"
)

func Showdown(someStruct any) (infoFields []string) {
	_type := reflect.TypeOf(someStruct)

	if _type.Kind() == reflect.Struct {
		numFields := _type.NumField()
		infoFields = make([]string, numFields)

		// Собираем поля структуры
		for i := 0; i < numFields; i++ {

			field := _type.Field(i)
			infoFields[i] = field.Name

			fmt.Println(field.Name)
		}
	}
	fmt.Println(infoFields)
	return infoFields
}

type Food struct {
	ID          int
	Name        string
	Type        string
	Category    string
	TotalPrice  float64
	Composition []string
}

func Cross(currStruct any, newMap map[string]interface{}) any {
	// Data of currStruct
	_type := reflect.TypeOf(currStruct).Elem()
	_value := reflect.ValueOf(currStruct).Elem()

	for i := 0; i < _type.NumField(); i++ {
		field := _type.Field(i)
		name := field.Name
		kind := field.Type.Kind()

		// If field is in Map for update, update
		if tmpValue, ok := newMap[name]; ok {
			newValue := reflect.ValueOf(tmpValue)
			updateField := _value.FieldByName(name)

			// Check in before wraiting
			if updateField.CanSet() && kind == newValue.Kind() {
				updateField.Set(newValue)
			}
		}
	}
	return currStruct
}

func main() {
	f := Food{ID: 1, Name: "Name1", Type: "Type1", Category: "Category1", TotalPrice: 10.25, Composition: []string{}}
	m := map[string]interface{}{"Name": "Name777"}

	f2 := Cross(&f, m)

	fmt.Println(f2)
}
