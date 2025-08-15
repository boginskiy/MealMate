package model

import (
	"reflect"
	"strings"
)

type Fooder interface {
	GetAttrs() []string
}

type Ingredient struct {
	ID             int
	Name           string
	Unit           string  // Unit is unit of measurement
	Quantity       float64 // Quantity is quantity ingredient in food
	CostOfUnit     float64 // CostOfUnit is cost of unit ingredient
	CostOfQuantity float64 // CostOfQuantity is all cost ingredient
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
