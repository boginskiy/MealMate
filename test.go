package main

import (
	"fmt"
	"reflect"
)

type Food struct {
	ID          int
	Name        string
	Type        string
	Category    string
	TotalPrice  float64
	Composition []string
}

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

func main() {
	f := Food{}
	Showdown(f)
}
