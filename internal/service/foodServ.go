package service

import (
	"encoding/json"
	"io"
	"mealmate/internal/db"
	m "mealmate/internal/model"
	"mealmate/pkg"
	"net/http"
)

type FoodServer interface {
	CreateFood(*http.Request) ([]byte, int)
	ReadFood(*http.Request) ([]byte, int)
}

type FoodServ struct {
	toolsStructer pkg.ToolsStructer
	dBer          db.DBFooder
	exFuncer      ExFuncer
}

func NewFoodServ(toolsS pkg.ToolsStructer, db db.DBFooder) *FoodServ {
	exFunc := NewExtraFunc()
	return &FoodServ{toolsStructer: toolsS, dBer: db, exFuncer: exFunc}
}

func (f *FoodServ) CreateFood(req *http.Request) ([]byte, int) {
	// Read Body
	tmpByte, err := io.ReadAll(req.Body)

	// Check error
	if err != nil {
		tmpErr := m.NewErrorWarn("error", http.StatusBadRequest, err.Error(), req.URL.Path)
		return tmpErr.PreparBody(req), http.StatusBadRequest
	}

	// Deserialization
	newFood := m.NewFood()
	err = json.Unmarshal(tmpByte, newFood)

	if err != nil {
		tmpErr := m.NewErrorWarn("error", http.StatusBadRequest, err.Error(), req.URL.Path)
		return tmpErr.PreparBody(req), http.StatusBadRequest
	}

	// Write in DB
	war := f.dBer.PutFood(newFood)
	if war != "" {
		tmpWar := m.NewErrorWarn("warning", http.StatusUnprocessableEntity, string(war), req.URL.Path)
		return tmpWar.PreparBody(req), http.StatusUnprocessableEntity
	}

	return tmpByte, http.StatusOK
}

/*
Params:

	?id=2
	&
	?name=Burger
	&
	?type=Vegetarian
	&
	?category=Fastfood
	&
	?totalPrice=10.25

	default - view all food
	todo    - set up pagination
*/
func (f *FoodServ) ReadFood(req *http.Request) ([]byte, int) {
	tmpStore := make([]m.Food, 0, 10) // tmpStore
	queryParams := req.URL.Query()    // Params of URL

	for _, food := range f.dBer.TakeFoodStore() {
		// Filter about queryParams
		if ok := f.exFuncer.NeedShow(*food, queryParams); ok {
			tmpStore = append(tmpStore, *food)
		}
	}

	// Serialization
	tmpByte, err := json.Marshal(tmpStore)
	if err != nil {
		tmpErr := m.NewErrorWarn("error", http.StatusBadRequest, err.Error(), req.URL.Path)
		return tmpErr.PreparBody(req), http.StatusBadRequest
	}

	return tmpByte, http.StatusOK
}
