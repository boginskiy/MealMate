package service

import (
	"encoding/json"
	"io"
	"mealmate/internal/db"
	m "mealmate/internal/model"
	"mealmate/pkg"
	"net/http"
)

const ID = "Name"

type FoodServ struct {
	exReflecter pkg.ExReflecter
	dBer        db.DBFooder
	exFuncer    ExFuncer
}

func NewFoodServ(exref pkg.ExReflecter, db db.DBFooder) *FoodServ {
	exfunc := NewExtraFunc()
	return &FoodServ{exReflecter: exref, dBer: db, exFuncer: exfunc}
}

func (f *FoodServ) Create(req *http.Request) ([]byte, int) {
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
func (f *FoodServ) Read(req *http.Request) ([]byte, int) {
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

func (f *FoodServ) Update(req *http.Request) ([]byte, int) {
	// Read Body
	tmpByte, err := io.ReadAll(req.Body)

	// Check error
	if err != nil {
		tmpErr := m.NewErrorWarn("error", http.StatusBadRequest, err.Error(), req.URL.Path)
		return tmpErr.PreparBody(req), http.StatusBadRequest
	}

	// Deserialization to map
	var forUpFood map[string]any
	err = json.Unmarshal(tmpByte, &forUpFood)

	if err != nil {
		tmpErr := m.NewErrorWarn("error", http.StatusBadRequest, err.Error(), req.URL.Path)
		return tmpErr.PreparBody(req), http.StatusBadRequest
	}

	// Check that is ID, like 'Name'
	tmpId, ok := forUpFood[ID]
	id, ok2 := tmpId.(string)

	if !ok || !ok2 {
		tmpErr := m.NewErrorWarn("error", http.StatusBadRequest, badBodyErr.Error(), req.URL.Path)
		return tmpErr.PreparBody(req), http.StatusBadRequest
	}

	newFood, war := f.dBer.ChangeFood(id, forUpFood)
	if war != "" {
		tmpWar := m.NewErrorWarn("warning", http.StatusBadRequest, string(war), req.URL.Path)
		return tmpWar.PreparBody(req), http.StatusBadRequest
	}

	// Serialization
	tmpByte2, err := json.Marshal(newFood)
	if err != nil {
		tmpErr := m.NewErrorWarn("error", http.StatusBadRequest, err.Error(), req.URL.Path)
		return tmpErr.PreparBody(req), http.StatusBadRequest
	}

	return tmpByte2, http.StatusOK
}

func (f *FoodServ) Delete(req *http.Request) ([]byte, int) {
	return []byte{}, 0
}
