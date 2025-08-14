package model

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

// func (f *Food) GetAttrs() []string {

// }
