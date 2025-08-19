package service

import (
	"bytes"
	"encoding/json"
	"io"
	m "mealmate/internal/model"
	"net/http"
)

type IngredServ struct {
}

func NewIngredServ() *IngredServ {
	return &IngredServ{}
}

// type IngredModel struct {
// 	ID             int
// 	Name           string
// 	Unit           string  // Unit is unit of measurement // Now Show for client
// 	Quantity       float64 // Quantity is quantity ingredient in food
// 	CostOfUnit     float64 // CostOfUnit is cost of unit ingredient // Now Show for client
// 	CostOfQuantity float64 // CostOfQuantity is all cost ingredient
// }

func (i *IngredServ) Read(req *http.Request) ([]byte, int) {
	return []byte{}, 0
}

func (i *IngredServ) Create(req *http.Request) ([]byte, int) {
	// Считать тело запроса
	newIngred := m.NewIngredModel()

	// Буфер
	var outBuf bytes.Buffer

	io.Copy(&newIngred, req.Body)

	json.Unmarshal(req.Body, &newIngred)

	return []byte{}, 0
}

func (i *IngredServ) Update(req *http.Request) ([]byte, int) {
	return []byte{}, 0
}

func (i *IngredServ) Delete(req *http.Request) ([]byte, int) {
	return []byte{}, 0
}
