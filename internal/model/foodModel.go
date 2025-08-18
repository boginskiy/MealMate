package model

import (
	"reflect"
	"strings"
)

type Fooder interface {
	GetAttrs() []string
}

type Food struct {
	ID          int
	Name        string
	Type        string
	Category    string
	TotalPrice  float64
	Composition []Ingredient
}

func NewFood() *Food {
	return &Food{}
}

func (f Food) GetAttrs() []string {
	_type := reflect.TypeOf(f)
	numFields := _type.NumField()
	fieldsList := make([]string, numFields)
	// Собираем поля структуры
	for i := 0; i < numFields; i++ {
		field := _type.Field(i)
		fieldsList[i] = strings.ToLower(field.Name)
	}
	return fieldsList
}
