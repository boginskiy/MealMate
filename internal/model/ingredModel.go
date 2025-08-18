package model

type IngredModel struct {
	ID             int
	Name           string
	Unit           string  // Unit is unit of measurement
	Quantity       float64 // Quantity is quantity ingredient in food
	CostOfUnit     float64 // CostOfUnit is cost of unit ingredient
	CostOfQuantity float64 // CostOfQuantity is all cost ingredient
}

func NewIngredModel() *IngredModel {
	return &IngredModel{}
}

func (f IngredModel) GetName() string {
	return f.Name
}
